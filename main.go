package main

import (
	"danielherschel/home-recipe/pkg/initializer"
	"danielherschel/home-recipe/pkg/middleware"
	bookRepo "danielherschel/home-recipe/pkg/repository/book"
	"danielherschel/home-recipe/pkg/router"
	"os"
)

func main() {
	initializer.LoadEnvs()

	repo := bookRepo.NewInMemoryRepository()
	r := router.NewRouter(repo).
		AddMiddleware(middleware.DevAuthMiddleware()).
		Build()
	
	// If behind a trusted proxy, set trusted proxies here
	// router.SetTrustedProxies([]string{"<trusted_proxy_ip>"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
