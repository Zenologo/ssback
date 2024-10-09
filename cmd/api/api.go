package api

import (
	"database/sql"
	"net/http"
	"time"

	logger "ssproxy/back/internal/pkg"
	"ssproxy/back/services/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	current_time := time.Now().Local()
	fileName := "main-" + current_time.Format("2006-01-02") + ".log"
	logger.InitLog("../logs/" + fileName)

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	logger.Info.Println("already to run http server")

	return http.ListenAndServe(s.addr, router)
}
