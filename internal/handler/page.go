package handler

import (
	"github.com/a-h/templ"
	"net/http"
	"time"
)

type PageHandler struct {
	tracker Tracker
}

func NewPageHandler(tracker Tracker) *PageHandler {
	return &PageHandler{
		tracker: tracker,
	}
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		level, _ := h.tracker.GetLevel(time.Now())
		component := page(level)
		layout := Layout(component, "Goffeine") //todo review
		templ.Handler(layout).ServeHTTP(w, r)
	case http.MethodPost:
		h.handlePagePost(w, r)
	default:
		http.Error(w, "Method not allowed, go away", http.StatusMethodNotAllowed)
	}
}

func (h *PageHandler) handlePagePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Form.Has("textinput") {
		input := r.Form.Get("textinput")

		err := h.tracker.Add(input)
		if err != nil {
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
