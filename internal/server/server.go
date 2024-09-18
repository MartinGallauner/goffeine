package server

import (
	"fmt"
	"net/http"
	"strings"
)

func GoffeineServer(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/status/")
	fmt.Fprint(w, GetStatus(user))
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
