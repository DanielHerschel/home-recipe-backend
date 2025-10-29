package main

import (
	"danielherschel/home-recipe/pkg/initializer"
	"danielherschel/home-recipe/pkg/middleware"
	bookRepo "danielherschel/home-recipe/pkg/repository/book"
	"danielherschel/home-recipe/pkg/router"
	"log"
	"os"
)

func main() {
	initializer.LoadEnvs()

	var repo bookRepo.RecipeBookRepository
	if os.Getenv("IN_MEMORY_DB") == "true" {
		log.Printf("Using in memory database")
		repo = bookRepo.NewInMemoryRepository()
	} else if url := os.Getenv("POSTGRES_URL"); url != "" {
		log.Printf("Using postgres database at %s", url)
		repo = bookRepo.NewRecipeBookRepository(url)
	} else {
		log.Fatal("No database url provided")
	}
	r := router.NewRouter(repo).
		AddMiddleware(middleware.DevAuthMiddleware()).
		Build()

	// If behind a trusted proxy, set trusted proxies here
	// router.SetTrustedProxies([]string{"<trusted_proxy_ip>"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
