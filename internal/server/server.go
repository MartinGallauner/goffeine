package server

import (
	"github.com/MartinGallauner/goffeine/internal/handler"
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

	handlers := handler.New(tracker)

	router := http.NewServeMux()
	router.Handle("/api/status", http.HandlerFunc(handlers.Status))
	router.Handle("/api/add", http.HandlerFunc(handlers.Intake))
	router.Handle("/", http.HandlerFunc(handlers.Page))

	fileServer := http.FileServer(http.Dir("./assets/dist"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	s.Handler = router
	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}
