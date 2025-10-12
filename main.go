package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/recipes/book/:name", func(c *gin.Context) {
		recipeBookname := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Get Recipe Book %s", recipeBookname)})
	})

	return router
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
