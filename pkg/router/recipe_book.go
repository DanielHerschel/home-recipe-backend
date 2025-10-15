package router

import (
	"danielherschel/home-recipe/pkg/domain"
	"danielherschel/home-recipe/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *Router) addRecipeBookRoutes() {
	router.GET("/api/books/:id", func(c *gin.Context) {
		getRecipeBook(c, router.RecipeBookService)
	})

	router.POST("/api/books/save", func(c *gin.Context) {
		saveRecipeBook(c, router.RecipeBookService)
	})

	router.GET("/api/books/list", func(c *gin.Context) {
		listRecipeBooks(c, router.RecipeBookService)
	})

	router.DELETE("/api/books/delete/:id", func(c *gin.Context) {
		deleteRecipeBook(c, router.RecipeBookService)
	})
}

func getRecipeBook(c *gin.Context, svc service.RecipeBookService) {
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
}

func saveRecipeBook(c *gin.Context, svc service.RecipeBookService) {
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
}

func deleteRecipeBook(c *gin.Context, svc service.RecipeBookService) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	recipeBookID := c.Param("id")
	if err := svc.DeleteRecipeBook(c.Request.Context(), userID, recipeBookID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

func listRecipeBooks(c *gin.Context, svc service.RecipeBookService) {
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
}