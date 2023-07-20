package database

import (
	"errors"
	"strings"

	"github.com/vucchaid/event-driven-architecture/event-service/database/types/mongodb"
	"github.com/vucchaid/event-driven-architecture/event-service/model"
)

type DatabaseHandler interface {
	AddEvent(model.Event) ([]byte, error)
	GetEvent([]byte) (model.Event, error)
	GetEventByName(string) (model.Event, error)
	GetAllEvents() ([]model.Event, error)
}

// To switch databases on go
func NewDatabaseConnection(connection, database string) (DatabaseHandler, error) {
	switch strings.ToLower(database) {
	case "mongodb":
		return mongodb.NewMongoDBConnection(connection)
	}
	return nil, errors.New("no database type provided")
}
