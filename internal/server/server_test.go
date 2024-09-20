package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETStatusUser(t *testing.T) {
	server := &GoffeineServer{}

	t.Run("returns current caffeine level of user 1", func(t *testing.T) {
		request := newGetStatusRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "100mg")
	})

	t.Run("returns current caffeine level of user 2", func(t *testing.T) {
		request := newGetStatusRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "50mg")
	})
}

func newGetStatusRequest(userId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/status/%s", userId), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubStore struct {
	status map[string]int
}

func (s *StubStore) GetStatus(name string) int {
	score := s.status[name]
	return score
}
