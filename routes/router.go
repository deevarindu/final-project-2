package routes

import (
	"github.com/deevarindu/final-project-2/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	user   *controllers.UserController
}

func NewRouter(router *gin.Engine, user *controllers.UserController) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/users", r.user.GetUsers)
	r.router.GET("/users/:id", r.user.GetUser)
	r.router.POST("/users", r.user.CreateUser)
	r.router.PUT("/users/:id", r.user.UpdateUser)
	r.router.DELETE("/users/:id", r.user.DeleteUser)
	r.router.Run(port)
}
