package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/health", healthCheck)
	router.Run("localhost:8080") // attaches the router to an http.Server and start server

}

func healthCheck(c *gin.Context) {
	//	c.IndentedJSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})

}
