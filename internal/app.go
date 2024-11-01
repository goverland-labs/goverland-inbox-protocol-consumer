package internal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goverland-labs/goverland-platform-events/pkg/natsclient"
	"github.com/nats-io/nats.go"
	"github.com/s-larionov/process-manager"

	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/config"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/feed"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/rest"
	apihandlers "github.com/goverland-labs/goverland-inbox-protocol-consumer/internal/rest/handlers"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/pkg/health"
	"github.com/goverland-labs/goverland-inbox-protocol-consumer/pkg/prometheus"
)

type Application struct {
	sigChan <-chan os.Signal
	manager *process.Manager
	cfg     config.App
}

func NewApplication(cfg config.App) (*Application, error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	a := &Application{
		sigChan: sigChan,
		cfg:     cfg,
		manager: process.NewManager(),
	}

	err := a.bootstrap()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) Run() {
	a.manager.StartAll()
	a.registerShutdown()
}

func (a *Application) bootstrap() error {
	initializers := []func() error{
		// Init Dependencies
		a.initServices,

		// Init Workers: Application
		a.initRestAPI,

		// Init Workers: System
		a.initPrometheusWorker,
		a.initHealthWorker,
	}

	for _, initializer := range initializers {
		if err := initializer(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Application) initServices() error {
	// TODO

	return nil
}

func (a *Application) initRestAPI() error {
	nc, err := nats.Connect(
		a.cfg.Nats.URL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(a.cfg.Nats.MaxReconnects),
		nats.ReconnectWait(a.cfg.Nats.ReconnectTimeout),
	)
	if err != nil {
		return err
	}

	pb, err := natsclient.NewPublisher(nc)
	if err != nil {
		return err
	}

	feedService := feed.NewService(pb)

	handlers := []apihandlers.APIHandler{
		apihandlers.NewFeedHandler(feedService),
	}

	a.manager.AddWorker(process.NewServerWorker("webhook api", rest.NewRestServer(a.cfg.REST, handlers)))

	return nil
}

func (a *Application) initPrometheusWorker() error {
	srv := prometheus.NewServer(a.cfg.Prometheus.Listen, "/metrics")
	a.manager.AddWorker(process.NewServerWorker("prometheus", srv))

	return nil
}

func (a *Application) initHealthWorker() error {
	srv := health.NewHealthCheckServer(a.cfg.Health.Listen, "/status", health.DefaultHandler(a.manager))
	a.manager.AddWorker(process.NewServerWorker("health", srv))

	return nil
}

func (a *Application) registerShutdown() {
	go func(manager *process.Manager) {
		<-a.sigChan

		manager.StopAll()
	}(a.manager)

	a.manager.AwaitAll()
}
