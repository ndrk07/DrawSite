package apiserver

import (
	"drawserver/internal/app/handlers"
	"drawserver/internal/app/model"
	"drawserver/internal/app/store"
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	sessionName = "gopher"
)

var (
	errIncorrectNameOrPassword = errors.New("incorrect name or password")
)

type server struct {
	router       *http.ServeMux
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       http.NewServeMux(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configurateRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configurateRouter() {
	// s.router.HandleFunc("/signin", handlers.SignInPage)
	s.router.HandleFunc("/signup", s.handleUsersCreate())
	s.router.HandleFunc("/signin", s.handleSessionsCreate())
	s.router.HandleFunc("/", handlers.HandleMainMenu)
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		log.Fatal(err)
	}

	type request struct {
		Name           string
		Password       string
		RepeatPassword string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			req := &request{
				Name:           r.FormValue("name"),
				Password:       r.FormValue("pw"),
				RepeatPassword: r.FormValue("pw_ag"),
			}

			if req.Password != req.RepeatPassword {
				handlers.ErrorOutputTmpl(w, r, tmpl, errors.New("Password must be the same"))
				return
			}
			u := &model.User{
				Name:     req.Name,
				Password: req.Password,
			}

			if err := s.store.User().Create(u); err != nil {
				handlers.ErrorOutputTmpl(w, r, tmpl, err)
				return
			}
			u.Sanitize()
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			if err = tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Render error", http.StatusInternalServerError)
			}
		}
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	tmpl, err := template.ParseFiles("templates/signin.html")
	if err != nil {
		log.Fatal(err)
	}

	type request struct {
		Name     string
		Password string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			req := &request{
				Name:     r.FormValue("name"),
				Password: r.FormValue("pw"),
			}

			u, err := s.store.User().FindByName(req.Name)
			if err != nil || !u.ComparePassword(req.Password) {
				handlers.ErrorOutputTmpl(w, r, tmpl, errIncorrectNameOrPassword)
				return
			}

			session, err := s.sessionStore.Get(r, sessionName)
			if err != nil {
				http.Error(w, "Session Error", http.StatusInternalServerError)
				return
			}

			session.Values["user_id"] = u.Id
			if err := s.sessionStore.Save(r, w, session); err != nil {
				http.Error(w, "Session Save Error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)

		} else {
			if err = tmpl.Execute(w, nil); err != nil {
				http.Error(w, "Render error", http.StatusInternalServerError)
			}
		}
	}
}
