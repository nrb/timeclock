package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type server struct {
	host, port string
	router     *gin.Engine
	db         *sql.DB
}

func NewServer(host, port string, router *gin.Engine, db *sql.DB) *server {
	return &server{
		host:   host,
		port:   port,
		router: router,
		db:     db,
	}
}

func (s *server) Run() {
	if s.host != "" && s.port != "" {
		s.router.Run(s.host, s.port)
	} else {
		s.router.Run()
	}
}
