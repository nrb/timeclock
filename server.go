package main

import "github.com/gin-gonic/gin"

type server struct {
	host, port string
	router     *gin.Engine
	// DB
}

func NewServer(host, port string, router *gin.Engine) *server {
	return &server{
		host:   host,
		port:   port,
		router: router,
	}
}

func (s *server) Run() {
	if s.host != "" && s.port != "" {
		s.router.Run(s.host, s.port)
	} else {
		s.router.Run()
	}
}
