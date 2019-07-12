package internal

import (
	"github.com/agoda-com/samsahai/pkg/apis/env/v1beta1"
)

const (
	//ArgoWorkflowDeployEngine     string = "argo-workflow"
	FluxHelmOperatorDeployEngine string = "flux-helm"
	MockDeployEngine             string = "mock"
)

type DeployEngine interface {
	// GetName returns name of deploy engine
	GetName() string

	// Create creates environment
	Create(refName string, comp *Component, parentComp *Component, values map[string]interface{}) error

	// Delete deletes environment
	Delete(queue *v1beta1.Queue) error

	// IsReady checks the environment is ready to use or not
	IsReady(queue *v1beta1.Queue) (bool, error)

	// GetLabelSelector returns map of label for select the components that created by the engine
	GetLabelSelectors(refName string) map[string]string

	// IsMocked uses for skip some functions due to mock deploy
	//
	// Skipped function: WaitForComponentsCleaned
	IsMocked() bool
}
