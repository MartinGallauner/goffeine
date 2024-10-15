package server

import (
	"fmt"
	"io"
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
	router.Handle("/", http.HandlerFunc(s.htmlHandler))

	s.Handler = router

	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func (s *GoffeineServer) htmlHandler(w http.ResponseWriter, r *http.Request) {

	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Simple Go HTML Server</title>
	</head>
	<body>
		<h1>Hello from Go!</h1>
		<p>This is a simple HTML page served by a Go application.</p>
	</body>
	</html>
	`
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, html)

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