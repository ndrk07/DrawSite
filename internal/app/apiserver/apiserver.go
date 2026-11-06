package apiserver

import (
	"database/sql"
	"drawserver/internal/app/store/sqlstore"
	"net/http"

	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)
	srv.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	globalServer := srv.middlewareAuth(srv.router)

	return http.ListenAndServe(config.BindAddr, globalServer)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
