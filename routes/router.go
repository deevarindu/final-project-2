package routes

import (
	"github.com/deevarindu/final-project-2/httpserver/controllers"
	"github.com/deevarindu/final-project-2/httpserver/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router       *gin.Engine
	user         *controllers.UserController
	photo        *controllers.PhotoController
	comment      *controllers.CommentController
	social_media *controllers.SocialMediaController
}

func NewRouter(router *gin.Engine, user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, social_media *controllers.SocialMediaController) *Router {
	return &Router{
		router:       router,
		user:         user,
		photo:        photo,
		comment:      comment,
		social_media: social_media,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/users", r.user.GetUsers)
	r.router.GET("/users/profile", r.user.GetUser)

	userRouter := r.router.Group("/users")
	{
		userRouter.POST("/register", r.user.Register)
		userRouter.POST("/login", r.user.Login)
		userRouter.PUT("/:id", middleware.Auth, r.user.UpdateUser)
		userRouter.DELETE("/:id", middleware.Auth, r.user.DeleteUser)
	}

	photoRouter := r.router.Group("/photos")
	{
		photoRouter.GET("/", r.photo.GetPhotos)
		photoRouter.POST("/", middleware.Auth, r.photo.UploadPhoto)
		photoRouter.PUT("/:id", middleware.Auth, r.photo.UpdatePhoto)
		photoRouter.DELETE("/:id", middleware.Auth, r.photo.DeletePhoto)
	}

	commentRouter := r.router.Group("/comments")
	{
		commentRouter.GET("/", r.comment.GetComments)
		commentRouter.POST("/", middleware.Auth, r.comment.CreateComment)
		commentRouter.PUT("/:id", middleware.Auth, r.comment.UpdateComment)
		commentRouter.DELETE("/:id", middleware.Auth, r.comment.DeleteComment)
	}

	socialMediaRouter := r.router.Group("/social-media")
	{
		socialMediaRouter.GET("/", r.social_media.GetSocialMedias)
		socialMediaRouter.POST("/", middleware.Auth, r.social_media.AddSocialMedia)
		socialMediaRouter.PUT("/:id", middleware.Auth, r.social_media.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middleware.Auth, r.social_media.DeleteSocialMedia)
	}

	r.router.Run(port)
}
