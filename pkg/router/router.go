package router

import (
	"danielherschel/home-recipe/pkg/middleware"
	"danielherschel/home-recipe/pkg/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	RecipeBookService service.RecipeBookService
}

func NewRouter(svc service.RecipeBookService) *Router {
	router := &Router{
		Engine:            setupRouterEngine(),
		RecipeBookService: svc,
	}

	router.addRecipeBookRoutes()
	return router
}

func setupRouterEngine() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.DevAuthMiddleware())

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return router
}

func getUserID(c *gin.Context) (string, bool) {
	uidIfc, ok := c.Get("userID")
	if !ok {
		c.JSON(401, gin.H{"error": "unauthenticated"})
		return "", false
	}
	userID := uidIfc.(string)
	return userID, true
}