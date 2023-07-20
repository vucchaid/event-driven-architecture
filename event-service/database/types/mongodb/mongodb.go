package mongodb

import (
	"github.com/vucchaid/event-service/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB     = "eventsDb"
	USERS  = "users"
	Events = "events"
)

type MongoDbLayer struct {
	session *mgo.Session
}

func NewMongoDBConnection(connection string) (*MongoDbLayer, error) {
	_, err := mgo.ParseURL(connection)
	if err != nil {
		return nil, err
	}
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDbLayer{
		session: session,
	}, nil
}

// Adds event to db
func (db *MongoDbLayer) AddEvent(event model.Event) ([]byte, error) {
	session := db.getFreshSession()
	defer session.Close()
	if !event.ID.Valid() {
		event.ID = bson.NewObjectId()
	}
	if !event.Location.ID.Valid() {
		event.Location.ID = bson.NewObjectId()
	}
	return []byte(event.ID), session.DB(DB).C(Events).Insert(event)
}

// Get event from db
func (db *MongoDbLayer) GetEvent(Id []byte) (model.Event, error) {
	session := db.getFreshSession()
	defer session.Close()
	event := model.Event{}
	err := session.DB(DB).C(Events).FindId(bson.ObjectId(Id)).One(&event)
	return event, err
}

// Get event by name from db
func (db *MongoDbLayer) GetEventByName(name string) (model.Event, error) {
	session := db.getFreshSession()
	defer session.Close()
	event := model.Event{}
	err := session.DB(DB).C(Events).Find(bson.M{"name": name}).One(&event)
	return event, err
}

// Get events from db
func (db *MongoDbLayer) GetAllEvents() ([]model.Event, error) {
	session := db.getFreshSession()
	defer session.Close()
	events := make([]model.Event, 0)
	err := session.DB(DB).C(Events).Find(nil).All(&events)
	return events, err
}

func (db *MongoDbLayer) getFreshSession() *mgo.Session {
	return db.session.Copy()
}
