package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter04/lib/msgqueue"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter04/lib/persistence"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
