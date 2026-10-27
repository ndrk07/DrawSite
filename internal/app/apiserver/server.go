package apiserver

import (
	"drawserver/internal/app/handlers"
	"drawserver/internal/app/model"
	"drawserver/internal/app/store"
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
	s.router.HandleFunc("/signup", handlers.SignUpPage)
	s.router.HandleFunc("/signin", handlers.SignInPage)
	s.router.HandleFunc("/postsignup", s.handleUsersCreate())
	s.router.HandleFunc("/", handlers.HandleMainMenu)
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Name           string
		Password       string
		RepeatPassword string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{
			Name:           r.FormValue("name"),
			Password:       r.FormValue("pw"),
			RepeatPassword: r.FormValue("pw_ag"),
		}

		if req.Password != req.RepeatPassword {
			handlers.ErrorSignUp(w, r, "Password must be the same")
			return
		}
		u := &model.User{
			Name:     req.Name,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {

			return
		}
		u.Sanitize()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
