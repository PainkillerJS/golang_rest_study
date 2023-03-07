package main

import (
	"log"
	"tutorial/todo-list"
	"tutorial/todo-list/pkg/handler"
	"tutorial/todo-list/pkg/repository"
	"tutorial/todo-list/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

}
