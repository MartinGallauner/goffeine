package handler

import (
	"net/http"
	"time"
)

type Handlers struct {
	statusHandler *StatusHandler
	intakeHandler *IntakeHandler
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func New(tracker Tracker) *Handlers {
	return &Handlers{
		statusHandler: NewStatusHandler(tracker),
		intakeHandler: NewIntakeHandler(tracker),
	}
}

func (h *Handlers) Status(w http.ResponseWriter, r *http.Request) {
	h.statusHandler.ServeHTTP(w, r)
}

func (h *Handlers) Intake(w http.ResponseWriter, r *http.Request) {
	h.intakeHandler.ServeHTTP(w, r)
}
