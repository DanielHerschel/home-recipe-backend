package router

import (
	"danielherschel/home-recipe/pkg/domain"
	"danielherschel/home-recipe/pkg/middleware"
	"danielherschel/home-recipe/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupRouter(svc service.RecipeBookService) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.DevAuthMiddleware())

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/api/books/:id", func(c *gin.Context) {
		userID, ok := getUserID(c)
        if !ok {
            return
        }

		recipeBookname := c.Param("id")
		recipeBook, err := svc.GetRecipeBook(c.Request.Context(), userID, recipeBookname)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if recipeBook == nil {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}
		c.JSON(200, recipeBook)
	})

	router.POST("/api/books/add", func(c *gin.Context) {
		userID, ok := getUserID(c)
        if !ok {
            return
        }

		recipeBook := &domain.RecipeBook{
			ID:     uuid.New().String(),
			Title:  "",
			UserID: "",
		}
		if err := c.ShouldBindJSON(recipeBook); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := svc.SaveRecipeBook(c.Request.Context(), userID, recipeBook); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "saved"})
	})

	router.GET("/api/books/list", func(c *gin.Context) {
		userID, ok := getUserID(c)
        if !ok {
            return
        }
		
		recipeBooks, err := svc.ListRecipeBooks(c.Request.Context(), userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, recipeBooks)
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