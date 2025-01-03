package server

import (
	"github.com/MartinGallauner/goffeine/internal/tracker"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGETStatusUser(t *testing.T) {
	server := NewGoffeineServer(&StubTracker{entries: make([]tracker.Entry, 0)}, &StubSessionManager{})

	t.Run("returns current caffeine level of user", func(t *testing.T) {
		request := newGetStatusRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Get status returns %v but expected 200.", response.Code)
		}
		assertResponseBody(t, response.Body.String(), "100")
	})
}

func newGetStatusRequest(userId string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/api/status", nil)
	return req
}

func TestPOSTAdd(t *testing.T) {
	server := NewGoffeineServer(&StubTracker{entries: make([]tracker.Entry, 0)}, &StubSessionManager{})

	t.Run("Adds one consumption of caffeine", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodPost, "/api/add", strings.NewReader("I had one cup of coffee a minute ago"))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusAccepted {
			t.Errorf("Get status returns %v but expected 202.", response.Code)
		}
	})
}

/*
func TestGETHome(t *testing.T) {
	server := NewGoffeineServer(&StubTracker{entries: make([]tracker.Entry, 0)})

	t.Run("Gets main page as html", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Returns status %v but expected 200", response.Code)
		}

		if !strings.Contains(response.Body.String(), "<html") {
			t.Error("Expected the <html tag in the response")

		}


	})

}
*/

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

type StubTracker struct {
	entries []tracker.Entry
}

type StubSessionManager struct {
}

func (s *StubSessionManager) LoadAndSave(handler http.Handler) http.Handler {
	return handler
}

func (s *StubTracker) GetLevel(time time.Time) (int, error) {
	return 100, nil
}

func (s *StubTracker) Add(userInput string) error {
	entry := tracker.Entry{Timestamp: time.Now(), CaffeineInMg: 100}
	s.entries = append(s.entries, entry)
	return nil
}
