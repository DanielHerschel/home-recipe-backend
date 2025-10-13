package main

import (
	"danielherschel/home-recipe/pkg/router"
	"danielherschel/home-recipe/pkg/service"
	"danielherschel/home-recipe/pkg/middleware"
)

func main() {
	svc := service.NewInMemoryService()
	r := router.NewRouter(svc).
		AddMiddleware(middleware.DevAuthMiddleware()).
		Build()

	r.Run(":8080")
}
