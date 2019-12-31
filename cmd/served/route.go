package served

import (
	"context"
	"einfach-msg/cmd/messaged"
	"einfach-msg/internal/kit/configuration"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Route struct{}

const (
	defaultPath = "./configurations/service.yaml"
)

func New() *Route {
	return &Route{}
}

func (r *Route) HTTP() {
	router := mux.NewRouter()
	config, err := configuration.Read(defaultPath)
	if err != nil {
		log.Fatalf("error while load configurations, got: %v", err)
	}

	v1 := router.PathPrefix("/v1").Subrouter()

	message := messaged.New()

	v1.HandleFunc("/message", message.Get).Methods(http.MethodGet)
	v1.HandleFunc("/message", message.Create).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: router,

		WriteTimeout: config.Server.WriteTimeout * time.Second,
		ReadTimeout:  config.Server.ReadTimeout * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while listen and serve, got: %v", err)
		}
	}()

	processCH := make(chan struct{})
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		terminationSignal := <-done
		log.Infof("caught signal: %v", terminationSignal)

		processCH <- struct{}{}
		signal.Stop(done)
	}()

	log.Infof("server starting listen and serve at port: %s", config.Server.Port)

	<-processCH

	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), config.Server.ShutdownTimeout*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error while shutdown the server, got: %v", err)
	}

	log.Info("server gracefully shutdown")
}
