package router

import (
	"mygram/config"
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	route := &controllers.Database{Connect: db}

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", route.UserRegister)
		userRouter.POST("/login", route.UserLogin)

		userRouter.PUT("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), route.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), middlewares.UserAuthorization(), route.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.POST("/", middlewares.Authentication(), route.AddPhoto)
		photoRouter.GET("/", middlewares.Authentication(), route.GetAllPhotos)
		photoRouter.PUT("/:photoId", middlewares.Authentication(), middlewares.PhotoAuthorization(), route.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.Authentication(), middlewares.PhotoAuthorization(), route.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", middlewares.Authentication(), route.AddComment)
		commentRouter.GET("/", middlewares.Authentication(), route.GetAllComments)
		commentRouter.PUT("/:commentId", middlewares.Authentication(), middlewares.CommentAuthorization(), route.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.Authentication(), middlewares.CommentAuthorization(), route.DeleteComment)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.POST("/", middlewares.Authentication(), route.CreateSosmed)
		sosmedRouter.GET("/", middlewares.Authentication(), route.GetSosmed)
		sosmedRouter.PUT("/:socialMediaId", middlewares.Authentication(), middlewares.SocmedAuthorization(), route.UpdateSosmed)
		sosmedRouter.DELETE("/:socialMediaId", middlewares.Authentication(), middlewares.SocmedAuthorization(), route.DeleteSosmed)
	}

	return r
}
