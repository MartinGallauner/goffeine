package server

import (
	"fmt"
	"net/http"
	"time"
)

type GoffeineServer struct {
	Tracker Tracker
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func (s *GoffeineServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//user := strings.TrimPrefix(r.URL.Path, "/level")
	//fmt.Fprint(w, s.Store.GetStatus(user))

	if r.URL.Path == "/status" {
		level, _ := s.Tracker.GetLevel(time.Now()) //todo handle error
		fmt.Fprint(w, level)
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, nil)
		return
	}

}

type CaffeineStore interface {
	GetStatus(user string) int
}
