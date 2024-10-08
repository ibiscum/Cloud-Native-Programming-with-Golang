package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter05/lib/msgqueue"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter05/lib/persistence"
)

func ServeAPI(endpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) error {
	handler := newEventHandler(dbHandler, eventEmitter)

	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.findEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)

	locationRouter := r.PathPrefix("/locations").Subrouter()
	locationRouter.Methods("GET").Path("").HandlerFunc(handler.allLocationsHandler)
	locationRouter.Methods("POST").Path("").HandlerFunc(handler.newLocationHandler)

	server := handlers.CORS()(r)

	return http.ListenAndServe(endpoint, server)
}
