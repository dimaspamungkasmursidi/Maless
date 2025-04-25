package routers

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/handlers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupROuter(middlewares middlewares.AuthMiddleware, oauthHandlers handlers.OauthHandlers, tagHandlers handlers.TagHandlers) *gin.Engine {

	r := gin.Default()

	oauth := r.Group("/oauth")
	oauth.GET("/auth/google/login", oauthHandlers.HandleGoogleLogin)
	oauth.GET("/callback", oauthHandlers.HandleGoogleCallback)

	// user := r.Group("/user")
	// user.POST("/create", UserHandlers.CreateUser)

	authGroup := r.Group("/tag")
	authGroup.Use(middlewares.ValidateUser()) // <- Middleware dipasang di sini
	{
		authGroup.POST("/create", tagHandlers.CreateTagHandler)
		authGroup.GET("/:idTag", tagHandlers.GetTagByIdHandler)
		authGroup.PUT("/update/:idTag", tagHandlers.UpdateTagHandler)
	}
	return r
}
