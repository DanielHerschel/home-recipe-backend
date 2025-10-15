package main

import (
	"danielherschel/home-recipe/pkg/router"
	repo "danielherschel/home-recipe/pkg/repository"
	"danielherschel/home-recipe/pkg/middleware"
)

func main() {
	svc := repo.NewInMemoryRepository()
	r := router.NewRouter(svc).
		AddMiddleware(middleware.DevAuthMiddleware()).
		Build()

	r.Run(":8080")
}
