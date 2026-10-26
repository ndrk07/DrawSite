package apiserver

import (
	"drawserver/internal/app/handlers"
	"drawserver/internal/app/store"
	"log/slog"
	"net/http"
	"os"
)

type APIserver struct {
	BindAddr string
	StoreURL string
	logger   *slog.Logger
	router   *http.ServeMux
	store    *store.Store
}

func New() *APIserver {
	return &APIserver{
		BindAddr: ":8080",
		StoreURL: "host=localhost dbname=drawsite sslnode=disable",
		logger:   slog.New(slog.NewTextHandler(os.Stdout, nil)),
		router:   http.NewServeMux(),
	}
}

func (s *APIserver) Start() error {
	s.logger.Info("starting api server")

	s.configurateRouter()

	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := s.configureStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.BindAddr, s.router)
}

func (s *APIserver) configurateRouter() {
	s.router.HandleFunc("/signup", handlers.SignUpPage)
	s.router.HandleFunc("/signin", handlers.SignInPage)
}

func (s *APIserver) configureStore() error {
	st := store.New(s.StoreURL)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}
