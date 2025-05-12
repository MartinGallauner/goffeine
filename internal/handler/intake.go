package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type IntakeHandler struct {
	tracker Tracker
}

func NewIntakeHandler(tracker Tracker) *IntakeHandler {
	return &IntakeHandler{
		tracker: tracker,
	}
}

func (h *IntakeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusAccepted)

	err = h.tracker.Add(string(body))
	if err != nil {
		slog.Info("While adding to the tracker, the following error occurred", "error", err)
		return
	}
	fmt.Fprint(w, nil)
}
