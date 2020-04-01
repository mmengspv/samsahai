package errors

import (
	"errors"

	pkgerrors "github.com/pkg/errors"
)

const (
	ErrInternalError             = Error("internal error")
	ErrNotImplemented            = Error("not implemented")
	ErrDeployTimeout             = Error("deploy timeout")
	ErrTestTimeout               = Error("test timeout")
	ErrTestRunnerNotFound        = Error("test runner not found")
	ErrRequestTimeout            = Error("request timeout")
	ErrExecutionTimeout          = Error("execution timeout")
	ErrImageVersionNotFound      = Error("image version not found")
	ErrNoDesiredComponentVersion = Error("no desired component version")

	ErrTeamNotFound               = Error("team not found")
	ErrTeamNamespaceStillCreating = Error("still creating namespace")
	ErrTeamNamespaceStillExists   = Error("destroyed namespace still exists")

	ErrActivePromotionTimeout         = Error("active promotion timeout")
	ErrActiveDemotionTimeout          = Error("demoted active environment timeout")
	ErrRollbackActivePromotionTimeout = Error("rollback active promotion timeout")
	ErrEnsureNamespaceDestroyed       = Error("namespace has not been destroyed")
	ErrEnsureActiveDemoted            = Error("active environment has been being demoted")
	ErrEnsureActivePromoted           = Error("active environment has been being promoted")
	ErrEnsureComponentDeployed        = Error("components has been being deployed")
	ErrEnsureComponentTested          = Error("components has been being tested")
	ErrDeletingReleases               = Error("deleting releases")
	ErrForceDeletingComponents        = Error("force deleting components")
	ErrRollingBackActivePromotion     = Error("rolling back active promotion process")

	ErrUnauthorized      = Error("unauthorized")
	ErrAuthTokenNotFound = Error("auth token not found")
	ErrInvalidJSONData   = Error("invalid json data")
	ErrCannotMarshalJSON = Error("cannot marshal to json")
	ErrCannotMarshalYAML = Error("cannot marshal to yaml")

	ErrTestConfigurationNotFound = Error("test configuration not found")

	ErrEnsureConfigDestroyed = Error("config been being destroyed")
)

var (
	Is    = errors.Is
	New   = errors.New
	Cause = pkgerrors.Cause
	Wrap  = pkgerrors.Wrap
	Wrapf = pkgerrors.Wrapf
)

type Error string

// Error overrides error
func (e Error) Error() string { return string(e) }

func IsImageNotFound(err error) bool {
	return ErrImageVersionNotFound.Error() == err.Error()
}

// IsTeamNotFound checks team is exist
func IsTeamNotFound(err error) bool {
	return ErrTeamNotFound.Error() == err.Error()
}

// IsNamespaceStillCreating checks namespace is still creating
func IsNamespaceStillCreating(err error) bool {
	return ErrTeamNamespaceStillCreating.Error() == err.Error()
}

// IsNamespaceStillExists checks namespace still exists
func IsNamespaceStillExists(err error) bool {
	return ErrTeamNamespaceStillExists.Error() == err.Error()
}

// IsEnsuringActivePromoted checks ensuring active promoted
func IsEnsuringActivePromoted(err error) bool {
	return ErrEnsureActivePromoted.Error() == err.Error()
}

// ErrEnsureActiveDemoted checks ensuring active demoted
func IsEnsuringActiveDemoted(err error) bool {
	return ErrEnsureActiveDemoted.Error() == err.Error()
}

// IsEnsuringComponentDeployed checks ensuring active deployed
func IsEnsuringComponentDeployed(err error) bool {
	return ErrEnsureComponentDeployed.Error() == err.Error()
}

// IsEnsuringActivePromoted checks ensuring active tested
func IsEnsuringActiveTested(err error) bool {
	return ErrEnsureComponentTested.Error() == err.Error()
}

// IsEnsuringNamespaceDestroyed checks ensuring namespace destroyed
func IsEnsuringNamespaceDestroyed(err error) bool {
	return ErrEnsureNamespaceDestroyed.Error() == err.Error()
}

// IsErrRequestTimeout checks request timeout
func IsErrRequestTimeout(err error) bool {
	return ErrRequestTimeout.Error() == err.Error()
}

// IsDeletingReleases checks releases have been deleting
func IsDeletingReleases(err error) bool {
	return ErrDeletingReleases.Error() == err.Error()
}

// IsErrActivePromotionTimeout checks active promotion has been timeout
func IsErrActivePromotionTimeout(err error) bool {
	return ErrActivePromotionTimeout.Error() == err.Error()
}

// IsErrActiveDemotionTimeout checks active demotion has been timeout
func IsErrActiveDemotionTimeout(err error) bool {
	return ErrActiveDemotionTimeout.Error() == err.Error()
}

// IsErrRollbackActivePromotionTimeout checks active promotion rollback has been timeout
func IsErrRollbackActivePromotionTimeout(err error) bool {
	return ErrRollbackActivePromotionTimeout.Error() == err.Error()
}

// IsRollingBackActivePromotion checks active promotion is rolling back
func IsRollingBackActivePromotion(err error) bool {
	return ErrRollingBackActivePromotion.Error() == err.Error()
}

// IsEnsuringConfigDestroyed checks ensuring config destroyed
func IsEnsuringConfigDestroyed(err error) bool {
	return ErrEnsureConfigDestroyed.Error() == err.Error()
}
