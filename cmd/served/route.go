package served

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GustafPahlevi/go-simple-svc/configurations"
	"github.com/GustafPahlevi/go-simple-svc/internal/app/message"
	"github.com/GustafPahlevi/go-simple-svc/internal/kit/configuration"
	mongoKit "github.com/GustafPahlevi/go-simple-svc/internal/kit/mongo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Route responsible handle http serve and register routes
type Route struct{}

const (
	defaultPath = "./configurations/service.yaml"
)

// New will instantiate Route itself
func New() *Route {
	return &Route{}
}

// HTTP is function that will serve HTTP, instantiate all dependencies and register all routes
func (r *Route) HTTP() {
	router := mux.NewRouter()
	config, err := configuration.Read(defaultPath)
	if err != nil {
		log.Fatalf("error while load configurations, got: %v", err)
	}

	msg, err := r.initMessageClient(config)
	if err != nil {
		log.Fatalf("error while initiate message client, got: %v", err)
	}

	log.Info("successfully initiate all dependencies")

	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"status":"service is ok"}`))
		if err != nil {
			log.Warnf("error while check service healthy, got: %v", err)
		}
	}).Methods(http.MethodGet)

	v1.HandleFunc("/message", msg.Get).Methods(http.MethodGet)
	v1.HandleFunc("/message", msg.Create).Methods(http.MethodPost)

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

func (r *Route) initMessageClient(config configurations.Structure) (*message.Message, error) {
	mongoURI := "mongodb://" + config.Database.Host + ":" + config.Database.Port
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return &message.Message{}, err
	}

	mongoCtx, _ := context.WithTimeout(context.Background(), config.Database.Timeout*time.Second)

	err = client.Connect(mongoCtx)
	if err != nil {
		return &message.Message{}, err
	}

	err = client.Ping(mongoCtx, readpref.Primary())
	if err != nil {
		return &message.Message{}, err
	}

	collection := client.Database(config.Database.Name).Collection(config.Database.Collection)
	collectionKit := mongoKit.New(collection, config.Database.Timeout)

	return message.New(collectionKit), nil
}
