package handler

import (
	"fmt"
	"net/http"
	"time"
)

type StatusHandler struct {
	tracker Tracker
}

func NewStatusHandler(tracker Tracker) *StatusHandler {
	return &StatusHandler{
		tracker: tracker,
	}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	level, _ := h.tracker.GetLevel(time.Now()) // TODO handle error
	fmt.Fprint(w, level)                       /* #nosec */
}
