package routers

import (
	"github.com/Bobby-P-dev/todo-listgo.git/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupROuter() *gin.Engine {

	r := gin.Default()
	r.GET("/auth/google/login", handlers.HandleGoogleLogin)
	r.GET("/callback", handlers.HandleGoogleCallback)
	return r

}
