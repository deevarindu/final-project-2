package main

import (
	"github.com/deevarindu/final-project-2/config"
	"github.com/deevarindu/final-project-2/httpserver/controllers"
	"github.com/deevarindu/final-project-2/httpserver/repositories/gorm"
	"github.com/deevarindu/final-project-2/httpserver/services"
	"github.com/deevarindu/final-project-2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.CreateConnectionGORM()
	if err != nil {
		panic(err)
	}

	userRepository := gorm.NewUserRepository(db)
	userSvc := services.NewUserSvc(userRepository)
	userHandler := controllers.NewUserController(userSvc)

	router := gin.Default()

	app := routes.NewRouter(router, userHandler)
	app.Start(":5000")
}
