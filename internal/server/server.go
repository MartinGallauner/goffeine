package server

import (
	"fmt"
	"net/http"
)

func GoffeineServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "100mg")
}
