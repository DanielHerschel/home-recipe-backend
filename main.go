package main

import (
	"danielherschel/home-recipe/pkg/router"
	"danielherschel/home-recipe/pkg/service"
)

func main() {
	svc := service.NewInMemoryService()
	r := router.SetupRouter(svc)

	r.Run(":8080")
}
