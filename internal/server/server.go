package server

import (
	"fmt"
	"io"
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
		level, _ := s.Tracker.GetLevel(time.Now()) // TODO handle error
		fmt.Fprint(w, level)
		return
	}

	if r.Method == http.MethodPost {
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

}

type CaffeineStore interface {
	GetStatus(user string) int
}
