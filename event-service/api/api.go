package api

import (
	"fmt"
	"net/http"

	"github.com/vucchaid/event-driven-architecture/event-service/database"
	"github.com/vucchaid/event-driven-architecture/event-service/event"

	"github.com/gorilla/mux"
)

func Serve(port, dbConnection, dbType string) error {

	router := mux.NewRouter()

	// Get db connection
	databaseHandler, err := database.NewDatabaseConnection(dbConnection, dbType)
	if err != nil {
		return err
	}
	// Instantiate Event handler
	eventHandler := event.NewEventServiceHandler(databaseHandler)

	eventsRouter := router.PathPrefix("/events").Subrouter()
	// Get single event based on Id or Name
	eventsRouter.Methods(http.MethodGet).Path("/{criteria}/{search}").HandlerFunc(eventHandler.GetEvent)
	// Get all events
	eventsRouter.Methods(http.MethodGet).Path("").HandlerFunc(eventHandler.GetAllEvents)
	// Post new event
	eventsRouter.Methods(http.MethodPost).Path("").HandlerFunc(eventHandler.AddEvent)

	// Serve
	return http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
