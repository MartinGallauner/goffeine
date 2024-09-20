package server

import (
	"fmt"
	"net/http"
	"strings"
)

type GoffeineServer struct {
	store CaffeineStore
}

func (s *GoffeineServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/status/")
	fmt.Fprint(w, s.store.GetStatus(user))
}

type CaffeineStore interface {
	GetStatus(user string) int
}

func GetStatus(user string) string {
	if user == "1" {
		return "100mg"
	}

	if user == "2" {
		return "50mg"
	}
	return ""
}
