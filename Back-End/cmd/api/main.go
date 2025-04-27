package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/db"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/handlers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/helpers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/middlewares"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/oauth"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/repository"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/routers"
	"github.com/Bobby-P-dev/todo-listgo.git/internal/services"
)

func init() {

}

func main() {
	helpers.LoadEnv()
	if err := oauth.InitOAuthConfig("http://localhost:8081/oauth/callback"); err != nil {
		log.Fatalf("Failed to initialize OAuth config: %v", err)
	}
	db := db.ConnectToDB()
	UserRepo := repository.NewUserRepoImpl(db)
	middleware := middlewares.NewAuthMiddleware(UserRepo)
	TagRepo := repository.NewTagRepoImpl(db)
	TagServices := services.NewTagServiceImpl(TagRepo)
	TagHandlers := handlers.NewTagHandlerImpl(TagServices)
	UserServices := services.NewUserServicesImpl(db, UserRepo)
	OauthHandlers := handlers.NewOauthHandlers(UserServices)

	r := routers.SetupROuter(*middleware, OauthHandlers, TagHandlers)
	s := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic("Server failed to start: " + err.Error())
	}
}
