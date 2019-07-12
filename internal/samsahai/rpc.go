package samsahai

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	s2h "github.com/agoda-com/samsahai/internal"
	s2herrors "github.com/agoda-com/samsahai/internal/errors"
	"github.com/agoda-com/samsahai/internal/reporter/slack"
	"github.com/agoda-com/samsahai/internal/samsahai/exporter"
	s2hv1beta1 "github.com/agoda-com/samsahai/pkg/apis/env/v1beta1"
	"github.com/agoda-com/samsahai/pkg/samsahai/rpc"
)

func (c *controller) authenticateRPC(authToken string) error {
	isMatch := subtle.ConstantTimeCompare([]byte(authToken), []byte(c.configs.SamsahaiCredential.InternalAuthToken))
	if isMatch != 1 {
		return s2herrors.ErrUnauthorized
	}
	return nil
}

func (c *controller) GetConfiguration(ctx context.Context, team *rpc.Team) (*rpc.Configuration, error) {
	if err := c.authenticateRPC(ctx.Value(s2h.SamsahaiAuthHeader).(string)); err != nil {
		return nil, err
	}

	teamName := team.Name
	cfg, ok := c.GetTeamConfigManager(teamName)
	if !ok {
		return nil, s2herrors.ErrTeamNotFound
	}
	data, err := json.Marshal(cfg.Get())
	if err != nil {
		return nil, errors.Wrapf(err, "cannot marshal configuration, team %s", teamName)
	}
	return &rpc.Configuration{Config: data, GitRevision: cfg.GetGitLatestRevision()}, nil
}

func (c *controller) GetMissingVersion(ctx context.Context, teamInfo *rpc.TeamWithCurrentComponent) (*rpc.ImageList, error) {
	if err := c.authenticateRPC(ctx.Value(s2h.SamsahaiAuthHeader).(string)); err != nil {
		return nil, err
	}

	configMgr, ok := c.GetTeamConfigManager(teamInfo.TeamName)
	if !ok {
		return nil, s2herrors.ErrTeamNotFound
	}

	teamComp := &s2hv1beta1.Team{}
	if err := c.getTeam(teamInfo.TeamName, teamComp); err != nil {
		return nil, errors.Wrapf(err, "cannot get of team %s", teamComp.Name)
	}

	stableList := &s2hv1beta1.StableComponentList{}

	if teamComp.Status.Namespace.Staging == "" {
		return nil, errors.Wrap(fmt.Errorf("staging namespace of %s is empty", teamInfo.TeamName),
			"staging namespace should not be empty")
	}

	err := c.client.List(ctx, &client.ListOptions{Namespace: teamComp.Status.Namespace.Staging}, stableList)
	if err != nil && !k8serrors.IsNotFound(err) {
		return nil, errors.Wrapf(err,
			"cannot get list of stable components, namespace %s", teamComp.Status.Namespace.Staging)
	}

	imgList := &rpc.ImageList{}
	comps := configMgr.GetComponents()
	for _, stable := range stableList.Items {
		source, ok := c.getImageSource(comps, stable.Name)
		if !ok {
			continue
		}

		// ignore current component
		if teamInfo.CompName == stable.Name {
			continue
		}

		c.detectAndAddImageMissing(*source, stable.Spec.Repository, stable.Name, stable.Spec.Version, imgList)
	}

	// add image missing for current component
	source, ok := c.getImageSource(comps, teamInfo.CompName)
	if ok {
		c.detectAndAddImageMissing(*source, teamInfo.Image.Repository, teamInfo.CompName, teamInfo.Image.Tag, imgList)
	}

	return imgList, nil
}

func (c *controller) detectAndAddImageMissing(source s2h.UpdatingSource, repo, name, version string, imgList *rpc.ImageList) {
	checker := c.checkers[string(source)]
	if err := checker.EnsureVersion(repo, name, version); err != nil {
		if s2herrors.IsImageNotFound(err) || s2herrors.IsErrRequestTimeout(err) {
			imgList.Images = append(imgList.Images, &rpc.Image{
				Repository: repo,
				Tag:        version,
			})
			return
		}
		logger.Error(err, "cannot ensure version",
			"name", name, "source", source, "repository", repo, "version", version)
	}
}

func (c *controller) getImageSource(comps map[string]*s2h.Component, name string) (*s2h.UpdatingSource, bool) {
	if _, ok := comps[name]; !ok {
		return nil, false
	}

	source := comps[name].Source
	if source == nil {
		return nil, false
	}
	if _, ok := c.checkers[string(*source)]; !ok {
		// ignore non-existing source
		return nil, false
	}

	return source, true
}

func (c *controller) NotifyComponentUpgrade(ctx context.Context, comp *rpc.ComponentUpgrade) (*rpc.Empty, error) {
	if err := c.authenticateRPC(ctx.Value(s2h.SamsahaiAuthHeader).(string)); err != nil {
		return nil, err
	}

	configMgr, ok := c.GetTeamConfigManager(comp.TeamName)
	if !ok {
		return &rpc.Empty{}, errors.Wrapf(s2herrors.ErrLoadConfiguration,
			"cannot load configuration, team %s", comp.TeamName)
	}

	queueHist := &s2hv1beta1.QueueHistory{}
	if err := c.getQueueHistory(comp.QueueHistoryName, comp.Namespace, queueHist); err != nil {
		return nil, errors.Wrapf(err,
			"cannot get queue history, name: %s, namespace: %s", comp.QueueHistoryName, comp.Namespace)
	}

	isReverify := queueHist.Spec.IsReverify
	isBuildSuccess := queueHist.Spec.IsTestSuccess && queueHist.Spec.IsDeploySuccess
	for _, reporter := range c.reporters {
		// TODO: user should be able to customize it via configuration
		// send slack notification if component upgrade failed
		if reporter.GetName() == slack.ReporterName {
			if comp.Status == rpc.ComponentUpgrade_SUCCESS {
				continue
			}
		}

		qHist := queueHist
		if comp.IssueType == rpc.ComponentUpgrade_DESIRED_VERSION_FAILED {
			var err error
			qHist, err = c.getLatestFailureQueueHistory(comp)
			if err != nil {
				return nil, err
			}
		}

		qHistName := qHist.Name
		testRunner := qHist.Spec.Queue.Status.TestRunner
		upgradeComp := s2h.NewComponentUpgradeReporter(
			comp,
			c.configs,
			s2h.WithIsReverify(isReverify),
			s2h.WithIsBuildSuccess(isBuildSuccess),
			s2h.WithTestRunner(testRunner),
			s2h.WithQueueHistoryName(qHistName),
		)
		c.sendComponentUpgradeReport(configMgr, reporter, upgradeComp)
	}

	if comp.Status == rpc.ComponentUpgrade_SUCCESS {
		if err := c.storeStableComponentsToTeam(ctx, comp); err != nil {
			return nil, err
		}
	}

	// Add metric updateQueueMetric & histories
	queueList := &s2hv1beta1.QueueList{}
	if err := c.client.List(context.TODO(), nil, queueList); err != nil {
		logger.Error(err, "cannot list all queue")
	}
	exporter.SetQueueMetric(queueList)

	queueHistoriesList := &s2hv1beta1.QueueHistoryList{}
	if err := c.client.List(context.TODO(), nil, queueHistoriesList); err != nil {
		logger.Error(err, "cannot list all queue")
	}
	exporter.SetQueueHistoriesMetric(queueHistoriesList, c.configs.SamsahaiURL)

	return &rpc.Empty{}, nil
}

func (c *controller) sendComponentUpgradeReport(configMgr s2h.ConfigManager, r s2h.Reporter, comp *s2h.ComponentUpgradeReporter) {
	if err := r.SendComponentUpgrade(configMgr, comp); err != nil {
		logger.Error(err, "cannot send component upgrade failure report", "team", comp.TeamName)
	}
}

func (c *controller) storeStableComponentsToTeam(ctx context.Context, comp *rpc.ComponentUpgrade) error {
	stableList := &s2hv1beta1.StableComponentList{}
	if err := c.client.List(ctx, &client.ListOptions{Namespace: comp.Namespace}, stableList); err != nil {
		return errors.Wrapf(err, "cannot get list of stable components, namespace %s", comp.Namespace)
	}

	teamComp := &s2hv1beta1.Team{}
	if err := c.getTeam(comp.TeamName, teamComp); err != nil {
		return errors.Wrapf(err, "cannot get of team %s", comp.TeamName)
	}

	teamComp.Status.SetStableComponents(stableList.Items)
	if err := c.updateTeam(teamComp); err != nil {
		return errors.Wrapf(err, "cannot update team %s", comp.TeamName)
	}

	return nil
}

func (c *controller) getQueueHistory(queueHistName, ns string, queueHist *s2hv1beta1.QueueHistory) error {
	return c.client.Get(context.TODO(), types.NamespacedName{Name: queueHistName, Namespace: ns}, queueHist)
}

func (c *controller) getLatestFailureQueueHistory(comp *rpc.ComponentUpgrade) (*s2hv1beta1.QueueHistory, error) {
	qLabels := s2h.GetDefaultLabels(comp.TeamName)
	qLabels["app"] = comp.Name
	qHists, err := c.listQueueHistory(qLabels)
	if err != nil {
		return &s2hv1beta1.QueueHistory{}, errors.Wrapf(err,
			"cannot list queue history, labels: %+v, namespace: %s", qLabels, comp.Namespace)
	}

	qHist := &s2hv1beta1.QueueHistory{}
	if len(qHists.Items) > 1 {
		qHist = &qHists.Items[1]
	}
	return qHist, nil
}

func (c *controller) listQueueHistory(selectors map[string]string) (*s2hv1beta1.QueueHistoryList, error) {
	queueHists := &s2hv1beta1.QueueHistoryList{}
	listOpt := &client.ListOptions{LabelSelector: labels.SelectorFromSet(selectors)}
	err := c.client.List(context.TODO(), listOpt, queueHists)
	queueHists.SortDESC()
	return queueHists, err
}
