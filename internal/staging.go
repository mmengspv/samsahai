package internal

import (
	"net/http"

	"github.com/agoda-com/samsahai/pkg/apis/env/v1beta1"
	stagingrpc "github.com/agoda-com/samsahai/pkg/staging/rpc"
)

type StagingTestRunner interface {
	// GetName returns type of test runner
	GetName() string

	// Trigger makes http request to run the test build
	Trigger(testConfig *ConfigTestRunner, currentQueue *v1beta1.Queue) error

	// GetResult makes http request to get result of test build [FAILURE/SUCCESS/UNKNOWN]
	// It returns bool results of is build success and is build finished
	GetResult(testConfig *ConfigTestRunner, currentQueue *v1beta1.Queue) (bool, bool, error)
}

type StagingController interface {
	// should implement RPC
	stagingrpc.RPC

	// should be able to serve http
	http.Handler

	// Start runs internal worker
	Start(stop <-chan struct{})

	// IsBusy returns true if controller still processing queue
	IsBusy() bool

	// LoadTestRunner loads single test runner to controller
	LoadTestRunner(runner StagingTestRunner)

	// LoadDeployEngine loads single deploy engine to controller
	LoadDeployEngine(engine DeployEngine)
}
