package exporter

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	s2hv1beta1 "github.com/agoda-com/samsahai/api/v1beta1"
	"github.com/agoda-com/samsahai/internal"
	s2hlog "github.com/agoda-com/samsahai/internal/log"
)

var logger = s2hlog.S2HLog.WithName("exporter")

var TeamMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "samsahai_team",
	Help: "List team name",
}, []string{"teamName"})

var HealthStatusMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "samsahai_health",
	Help: "show s2h's health status",
}, []string{"version", "gitCommit"})

var QueueMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "samsahai_queue",
	Help: "Show components in queue",
}, []string{"order", "teamName", "component", "version", "state", "no_of_processed"})

var ActivePromotionMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "samsahai_active_promotion",
	Help: "Get values from samsahai active promotion",
}, []string{"teamName", "state"})

func RegisterMetrics() {
	metrics.Registry.MustRegister(TeamMetric)
	metrics.Registry.MustRegister(QueueMetric)
	metrics.Registry.MustRegister(ActivePromotionMetric)
	metrics.Registry.MustRegister(HealthStatusMetric)
}

func SetTeamNameMetric(teamList map[string]internal.ConfigManager) {
	for teamName := range teamList {
		TeamMetric.WithLabelValues(teamName).Set(1)
	}
}

func SetQueueMetric(queue *s2hv1beta1.Queue) {
	queueStateList := map[string]float64{"waiting": 0, "testing": 0, "finished": 0, "deploying": 0, "cleaning": 0}
	switch queue.Status.State {
	case s2hv1beta1.Waiting:
		queueStateList["waiting"] = 1
		for state, val := range queueStateList {
			QueueMetric.WithLabelValues(
				strconv.Itoa(queue.Spec.NoOfOrder),
				queue.Spec.TeamName,
				queue.Name,
				queue.Spec.Version,
				state,
				strconv.Itoa(queue.Status.NoOfProcessed)).Set(val)
		}
	case s2hv1beta1.Testing, s2hv1beta1.Collecting:
		queueStateList["testing"] = 1
		for state, val := range queueStateList {
			QueueMetric.WithLabelValues(
				strconv.Itoa(queue.Spec.NoOfOrder),
				queue.Spec.TeamName,
				queue.Name,
				queue.Spec.Version,
				state,
				strconv.Itoa(queue.Status.NoOfProcessed)).Set(val)
		}
	case s2hv1beta1.Finished:
		queueStateList["finished"] = 1
		for state, val := range queueStateList {
			QueueMetric.WithLabelValues(
				strconv.Itoa(queue.Spec.NoOfOrder),
				queue.Spec.TeamName,
				queue.Name,
				queue.Spec.Version,
				state,
				strconv.Itoa(queue.Status.NoOfProcessed)).Set(val)
		}
	case s2hv1beta1.DetectingImageMissing, s2hv1beta1.Creating:
		queueStateList["deploying"] = 1
		for state, val := range queueStateList {
			QueueMetric.WithLabelValues(
				strconv.Itoa(queue.Spec.NoOfOrder),
				queue.Spec.TeamName,
				queue.Name,
				queue.Spec.Version,
				state,
				strconv.Itoa(queue.Status.NoOfProcessed)).Set(val)
		}
	case s2hv1beta1.CleaningBefore, s2hv1beta1.CleaningAfter:
		queueStateList["cleaning"] = 1
		for state, val := range queueStateList {
			QueueMetric.WithLabelValues(
				strconv.Itoa(queue.Spec.NoOfOrder),
				queue.Spec.TeamName,
				queue.Name,
				queue.Spec.Version,
				state,
				strconv.Itoa(queue.Status.NoOfProcessed)).Set(val)
		}
	}
}

func SetHealthStatusMetric(version, gitCommit string, ts float64) {
	HealthStatusMetric.WithLabelValues(
		version,
		gitCommit).Set(ts)
}

func SetActivePromotionMetric(activeProm *s2hv1beta1.ActivePromotion) {
	activePromStateList := map[string]float64{"waiting": 0, "deploying": 0, "testing": 0, "promoting": 0, "destroying": 0}
	atpState := activeProm.Status.State
	if atpState != "" {
		switch atpState {
		case s2hv1beta1.ActivePromotionWaiting:
			activePromStateList["waiting"] = 1
			for state, val := range activePromStateList {
				ActivePromotionMetric.WithLabelValues(
					activeProm.Name,
					state).Set(val)
			}
		case s2hv1beta1.ActivePromotionDeployingComponents, s2hv1beta1.ActivePromotionCreatingPreActive:
			activePromStateList["deploying"] = 1
			for state, val := range activePromStateList {
				ActivePromotionMetric.WithLabelValues(
					activeProm.Name,
					state).Set(val)
			}
		case s2hv1beta1.ActivePromotionTestingPreActive, s2hv1beta1.ActivePromotionCollectingPreActiveResult:
			activePromStateList["testing"] = 1
			for state, val := range activePromStateList {
				ActivePromotionMetric.WithLabelValues(
					activeProm.Name,
					state).Set(val)
			}
		case s2hv1beta1.ActivePromotionActiveEnvironment, s2hv1beta1.ActivePromotionDemoting:
			activePromStateList["promoting"] = 1
			for state, val := range activePromStateList {
				ActivePromotionMetric.WithLabelValues(
					activeProm.Name,
					state).Set(val)
			}
		case s2hv1beta1.ActivePromotionDestroyingPreActive, s2hv1beta1.ActivePromotionDestroyingPreviousActive,
			s2hv1beta1.ActivePromotionFinished:
			activePromStateList["destroying"] = 1
			for state, val := range activePromStateList {
				ActivePromotionMetric.WithLabelValues(
					activeProm.Name,
					state).Set(val)
			}

		}
	}
}
