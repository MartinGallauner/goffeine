package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStatusUser1(t *testing.T) {
	t.Run("returns current coffeine level of user 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/status/1", nil)
		response := httptest.NewRecorder()

		GoffeineServer(response, request)

		got := response.Body.String()
		want := "100mg"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGETStatusUser2(t *testing.T) {
	t.Run("returns current coffeine level of user 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/status/2", nil)
		response := httptest.NewRecorder()

		GoffeineServer(response, request)

		got := response.Body.String()
		want := "50mg"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
