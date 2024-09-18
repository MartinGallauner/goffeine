package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStatus(t *testing.T) {
	t.Run("returns current coffeine level", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/status", nil)
		response := httptest.NewRecorder()

		GoffeineServer(response, request)

		got := response.Body.String()
		want := "100mg"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
