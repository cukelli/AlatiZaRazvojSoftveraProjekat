package main

import (
	"context"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/config"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/metrics"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/poststore"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ps, err := poststore.New()
	if err != nil {
		log.Fatal(err)
	}

	service := &service.Service{
		Configurations: []*config.Config{},
		PostStore:      ps,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/configurations", metrics.Count(service.AddConfiguration, "/configurations")).Methods("POST")
	router.HandleFunc("/configurations/{id}/{version}", metrics.Count(service.GetConfiguration, "/configurations/{id}/{version}")).Methods("GET")
	router.HandleFunc("/configurations/{id}/{version}", metrics.Count(service.DeleteConfiguration, "/configurations/{id}/{version}")).Methods("DELETE")
	router.HandleFunc("/group", metrics.Count(service.AddConfigurationGroup, "/group")).Methods("POST")
	router.HandleFunc("/group/{id}/{version}", metrics.Count(service.GetConfigurationGroup, "/group/{id}/{version}")).Methods("GET")
	router.HandleFunc("/group/{id}/{version}", metrics.Count(service.DeleteConfigurationGroup, "/group/{id}/{version}")).Methods("DELETE")
	router.HandleFunc("/group/{id}/{version}/extend", metrics.Count(service.ExtendConfigurationGroup, "/group/{id}/{version}/extend")).Methods("POST")
	router.HandleFunc("/swagger.yaml", metrics.Count(service.SwaggerHandler, "/swagger.yaml")).Methods("GET")
	router.HandleFunc("/group/{id}/{version}/{labels}", metrics.Count(service.GetConfigurationGroupsByLabels, "/group/{id}/{version}/{labels}")).Methods("GET")

	// Prometheus metrics endpoint
	router.Handle("/metrics", promhttp.Handler())

	// start server
	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
