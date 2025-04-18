package server

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"io"
	"log"
	"net/http"
	"time"
)

type GoffeineServer struct {
	Tracker Tracker
	http.Handler
}

func NewGoffeineServer(tracker Tracker) *GoffeineServer {
	s := &GoffeineServer{
		Tracker: tracker,
	}

	router := http.NewServeMux()
	router.Handle("/api/status", http.HandlerFunc(s.statusHandler))
	router.Handle("/api/add", http.HandlerFunc(s.intakeHandler))
	router.Handle("/", http.HandlerFunc(s.handlePage))
	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

type SessionManager interface {
	LoadAndSave(http.Handler) http.Handler
	GetString(context.Context, string) string
	Put(context.Context, string, interface{})
}

func (s *GoffeineServer) handlePage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		level, _ := s.Tracker.GetLevel(time.Now())
		component := page(level)
		templ.Handler(component).ServeHTTP(w, r)
	case http.MethodPost:
		s.handlePagePost(w, r)
	default:
		http.Error(w, "Method not allowed, go away", http.StatusMethodNotAllowed)
	}
}

func (s *GoffeineServer) handlePagePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	if r.Form.Has("textinput") {
		input := r.Form.Get("textinput")

		err := s.Tracker.Add(input)
		if err != nil {
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *GoffeineServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	level, _ := s.Tracker.GetLevel(time.Now()) // TODO handle error
	fmt.Fprint(w, level)                       /* #nosec */
}

func (s *GoffeineServer) intakeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusAccepted)

	err = s.Tracker.Add(string(body))
	if err != nil {
		log.Printf("While adding to the tracker, the following error occured: %v", err)
		return
	}
	fmt.Fprint(w, nil)
}
