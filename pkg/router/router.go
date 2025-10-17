package router

import (
	bookRepo "danielherschel/home-recipe/pkg/repository/book"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
	RecipeBookService bookRepo.RecipeBookRepository
}

type RouterBuilder struct {
	router *Router
}

func NewRouter(repo bookRepo.RecipeBookRepository) *RouterBuilder {
	return &RouterBuilder{
		router: &Router{
			Engine:            gin.Default(),
			RecipeBookService: repo,
		},
	}
}

func (builder *RouterBuilder) AddMiddleware(mw gin.HandlerFunc) *RouterBuilder {
	builder.router.Use(mw)
	return builder
}

func (builder *RouterBuilder) Build() *Router {
	builder.router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	builder.router.addRecipeBookRoutes()

	return builder.router
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
