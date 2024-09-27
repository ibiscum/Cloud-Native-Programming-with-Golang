package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter06/lib/msgqueue"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter06/lib/persistence"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter, database})

	srv := http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	srv.ListenAndServe()
}