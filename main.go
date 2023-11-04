package main

import (
	"log"
	"net/http"
	v1 "sports-day/router/v1"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// check if location exists
	_, err := time.LoadLocation("Asia/Calcutta")
	if err != nil {

		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP) //can access ip address by r.RemoteAddr anywhere

	// simple health check on base url hit
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("health ok"))
	})

	r.Route("/v1", v1.Routes)

	log.Println("go-api server started")

	server := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: 3 * time.Minute,
		Handler:           r,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
