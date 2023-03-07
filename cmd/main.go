package main

import (
	"log"
	"tutorial/todo-list"
	"tutorial/todo-list/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handler := new(handler.Handler)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

}
