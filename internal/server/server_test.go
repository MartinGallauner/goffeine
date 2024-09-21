package server

import (
	"fmt"
	"github.com/MartinGallauner/goffeine/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGETStatusUser(t *testing.T) {
	server := &GoffeineServer{
		Tracker: &StubTracker{entries: make([]repository.Entry, 0)},
	}

	t.Run("returns current caffeine level of user", func(t *testing.T) {
		request := newGetStatusRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "100")
	})

}

func newGetStatusRequest(userId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/status"), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubTracker struct {
	entries []repository.Entry
}

func (s *StubTracker) GetLevel(time time.Time) (int, error) {
	return 100, nil
}

func (s *StubTracker) Add(userInput string) error {
	entry := repository.Entry{Timestamp: time.Now(), CaffeineInMg: 100}
	s.entries = append(s.entries, entry)
	return nil
}
