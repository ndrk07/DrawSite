package apiserver

import (
	"drawserver/internal/app/handlers"
	"drawserver/internal/app/model"
	"drawserver/internal/app/store"
	"errors"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type server struct {
	router *http.ServeMux
	logger *slog.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: http.NewServeMux(),
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
		store:  store,
	}

	s.configurateRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configurateRouter() {
	s.router.HandleFunc("/signin", handlers.SignInPage)
	s.router.HandleFunc("/signup", s.handleUsersCreate())
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
				handlers.ErrorSignUp(w, r, tmpl, errors.New("Password must be the same"))
				return
			}
			u := &model.User{
				Name:     req.Name,
				Password: req.Password,
			}

			if err := s.store.User().Create(u); err != nil {
				handlers.ErrorSignUp(w, r, tmpl, err)
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
