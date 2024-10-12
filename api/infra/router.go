package infra

import (
	"api/handler"
	"api/repository"
	"api/usecase"
	"log"
	"net/http"
)

func StartServer() {
	toDoRepo := repository.NewToDoRepositoryInMemory()
	toDoUsecase := usecase.NewToDoUsecase(toDoRepo)
	toDoHandler := handler.NewTodoHandler(toDoUsecase)

	http.HandleFunc("/todos", toDoHandler.GetAllToDos)
	http.HandleFunc("/todos/create", toDoHandler.CreateToDo)
	http.HandleFunc("/todos/delete", toDoHandler.DeleteToDo)

	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
