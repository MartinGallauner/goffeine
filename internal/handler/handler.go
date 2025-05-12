package handler

import (
	"net/http"
	"time"
)

type Handlers struct {
	statusHandler *StatusHandler
	intakeHandler *IntakeHandler
	pageHandler   *PageHandler
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func New(tracker Tracker) *Handlers {
	return &Handlers{
		statusHandler: NewStatusHandler(tracker),
		intakeHandler: NewIntakeHandler(tracker),
		pageHandler:   NewPageHandler(tracker),
	}
}

func (h *Handlers) Status(w http.ResponseWriter, r *http.Request) {
	h.statusHandler.ServeHTTP(w, r)
}

func (h *Handlers) Intake(w http.ResponseWriter, r *http.Request) {
	h.intakeHandler.ServeHTTP(w, r)
}

func (h *Handlers) Page(w http.ResponseWriter, r *http.Request) {
	h.pageHandler.ServeHTTP(w, r)
}
