package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/user", createUser)
	r.GET("/user/:id", getUser)
	r.PUT("/user/:id", updateUser)
	r.DELETE("/user/:id", deleteUser)
	r.GET("/users/", listUsers)

	r.POST("/entry", createEntry)
	r.GET("/entry/:id", getEntry)
	r.PUT("/entry/:id", updateEntry)
	r.DELETE("/entry/:id", deleteEntry)
	r.GET("/entries/", listEntries)

	// TODO: Admin routes

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func createUser(c *gin.Context) {
}

func getUser(c *gin.Context) {
}

func updateUser(c *gin.Context) {
}

func deleteUser(c *gin.Context) {
}

func listUsers(c *gin.Context) {
}

func createEntry(c *gin.Context) {
}

func getEntry(c *gin.Context) {
}

func updateEntry(c *gin.Context) {
}

func deleteEntry(c *gin.Context) {
}

func listEntries(c *gin.Context) {
}
