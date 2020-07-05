package workers

import (
	"github.com/labstack/echo"

	"github.com/gaia-pipeline/gaia/security"
	"github.com/gaia-pipeline/gaia/store"
	"github.com/gaia-pipeline/gaia/workers/scheduler/service"
)

// Dependencies define dependencies which this service needs.
type Dependencies struct {
	Scheduler   service.GaiaScheduler
	Certificate security.CAAPI
	Store       store.GaiaStore
}

// NewDependencies creates a new Dependencies struct.
func NewDependencies(scheduler service.GaiaScheduler, certificate security.CAAPI, store store.GaiaStore) Dependencies {
	return Dependencies{Scheduler: scheduler, Certificate: certificate, Store: store}
}

type workerProvider struct {
	deps Dependencies
}

// WorkerProvider defines functionality which this provider provides.
type WorkerProvider interface {
	RegisterWorker(c echo.Context) error
	DeregisterWorker(c echo.Context) error
	GetWorkerRegisterSecret(c echo.Context) error
	GetWorkerStatusOverview(c echo.Context) error
	ResetWorkerRegisterSecret(c echo.Context) error
	GetWorker(c echo.Context) error
}

// NewWorkerProvider creates a provider which provides worker related functionality.
func NewWorkerProvider(deps Dependencies) WorkerProvider {
	return &workerProvider{
		deps: deps,
	}
}
