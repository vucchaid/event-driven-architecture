package event

import (
	"encoding/hex"
	"encoding/json"
	"event-service/database"
	"event-service/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type eventService struct {
	dbHandler database.DatabaseHandler
}

func NewEventServiceHandler(dbHandler database.DatabaseHandler) *eventService {
	return &eventService{
		dbHandler: dbHandler,
	}
}

// Get event
func (service *eventService) GetEvent(res http.ResponseWriter, req *http.Request) {

	parameters := mux.Vars(req)

	criteria, ok := parameters["criteria"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		RespondErrorJSON(res, "No criteria found. It should be either name or id.")
		return
	}

	key, ok := parameters["search"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		RespondErrorJSON(res, "No criteria found. It should be either name or id.")
		return
	}

	var event model.Event
	var err error

	if strings.ToLower(criteria) == "name" {
		event, err = service.dbHandler.GetEventByName(key)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			RespondErrorJSON(res, err.Error())
			return
		}
	}

	if strings.ToLower(criteria) == "id" {
		id, err := hex.DecodeString(key)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			RespondErrorJSON(res, "No criteria found. It should be either name or id.")
			return
		}
		event, err = service.dbHandler.GetEvent(id)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			RespondErrorJSON(res, err.Error())
			return
		}
	}

	res.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(res).Encode(&event)
}

// Get all events
func (service *eventService) GetAllEvents(res http.ResponseWriter, req *http.Request) {
	events, err := service.dbHandler.GetAllEvents()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		RespondErrorJSON(res, err.Error())
		return
	}
	res.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(res).Encode(&events)
}

// Add an event
func (service *eventService) AddEvent(res http.ResponseWriter, req *http.Request) {
	event := model.Event{}
	err := json.NewDecoder(req.Body).Decode(&event)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		RespondErrorJSON(res, err.Error())
		return
	}
	_, err = service.dbHandler.AddEvent(event)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		RespondErrorJSON(res, err.Error())
		return
	}
	fmt.Fprint(res, `{"status":"success"}`)
}

func RespondErrorJSON(res http.ResponseWriter, message string) {
	message = fmt.Sprintf(`{"error":"%v"}`, message)
	fmt.Fprint(res, message)
}
