package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/a-h/templ"
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
	//router.Handle("/", http.HandlerFunc(s.htmlHandler))

	component := hello("galle")
	router.Handle("/", templ.Handler(component))

	s.Handler = router

	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func (s *GoffeineServer) htmlHandler(w http.ResponseWriter, r *http.Request) {

	component := hello("galle")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, component)

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
