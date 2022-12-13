package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/istomin10593/bookstore_items-api/src/clients/elasticsearch"
	"github.com/istomin10593/bookstore_utils-go/logger"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8083",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	logger.Info("about to start the application")
}
