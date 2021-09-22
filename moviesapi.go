package main

import (
	"net/http"
	"time"

	"github.com/un1xcorn/moviesapi/api"
	"github.com/un1xcorn/moviesapi/config"
	"github.com/un1xcorn/moviesapi/router"

	log "github.com/sirupsen/logrus"
)

func main() {

	api.Init()
	r := router.Init()
	c := config.LoadOrDefault("config.json")
	srv := &http.Server{
		Handler:      r,
		Addr:         c.Host + ":" + c.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server started (", c.Host+":"+c.Port, ")")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("An error happened while starting http server.")
	}

}
