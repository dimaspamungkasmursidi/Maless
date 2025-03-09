package main

import (
	"net/http"
	"time"

	"github.com/Bobby-P-dev/todo-listgo.git/internal/routers"
)

func init() {

}

func main() {

	r := routers.SetupROuter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic("Server failed to start: " + err.Error())
	}
}
