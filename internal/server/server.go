package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type GoffeineServer struct {
	Tracker Tracker
	Router *http.ServeMux
}

func NewGoffeineServer(tracker Tracker) *GoffeineServer {
	s := &GoffeineServer{
		Tracker: tracker,
		Router: http.NewServeMux(),
	}

	s.Router.Handle("/api/status", http.HandlerFunc(s.statusHandler))
	s.Router.Handle("/api/add", http.HandlerFunc(s.intakeHandler))

	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func (s *GoffeineServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w,r)
}

func (s *GoffeineServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	level, _ := s.Tracker.GetLevel(time.Now()) // TODO handle error
		fmt.Fprint(w, level)
		return
}

func (s *GoffeineServer) intakeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		w.WriteHeader(http.StatusAccepted)

		s.Tracker.Add(string(body))
		fmt.Fprint(w, nil)
		return
}