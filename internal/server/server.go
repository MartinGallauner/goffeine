package server

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
)

type GoffeineServer struct {
	Tracker Tracker
	http.Handler
	SessionManager *scs.SessionManager
}

func NewGoffeineServer(tracker Tracker, sessionManager *scs.SessionManager) *GoffeineServer {
	s := &GoffeineServer{
		Tracker: tracker,
	}

	router := http.NewServeMux()
	router.Handle("/api/status", http.HandlerFunc(s.statusHandler))
	router.Handle("/api/add", http.HandlerFunc(s.intakeHandler))
	router.Handle("/", http.HandlerFunc(s.handlePage))

	routerWithMiddleware := sessionManager.LoadAndSave(router)
	
	s.Handler = routerWithMiddleware
	return s
}

type Tracker interface {
	GetLevel(time time.Time) (int, error)
	Add(userInput string) error
}

func (s *GoffeineServer) handlePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		level,_ := s.Tracker.GetLevel(time.Now())
		component := page(level)
		templ.Handler(component).ServeHTTP(w, r)
	case http.MethodPost:
		s.handlePagePost(w, r)		
	default:
		http.Error(w, "Method not allowed, go away", http.StatusMethodNotAllowed)	
	}
}

func (s *GoffeineServer) handlePagePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Has("textinput") {
		input := r.Form.Get("textinput")
		s.Tracker.Add(input)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func (s *GoffeineServer) statusHandler(w http.ResponseWriter, r *http.Request) {
	level, _ := s.Tracker.GetLevel(time.Now()) // TODO handle error
	fmt.Fprint(w, level)
	return
}

func (s *GoffeineServer) intakeHandler(w http.ResponseWriter, r *http.Request) {
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
