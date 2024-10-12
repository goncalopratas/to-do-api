package repository

import (
	"api/domain"
	"errors"
)

type ToDoRepositoryInMemory struct {
	toDos []domain.ToDo
}

func NewToDoRepositoryInMemory() *ToDoRepositoryInMemory {
	return &ToDoRepositoryInMemory{
		toDos: []domain.ToDo{
			{ID: "1", Title: "Buy Birthday Gift", Description: "Dont forget GF birthday, buy something !", Status: domain.Pending},
			{ID: "2", Title: "Buy Chocapic", Description: "CHOCAPIC !", Status: domain.Completed},
		},
	}
}

func (r *ToDoRepositoryInMemory) GetAll() ([]domain.ToDo, error) {
	return r.toDos, nil
}

func (r *ToDoRepositoryInMemory) Add(t domain.ToDo) error {
	r.toDos = append(r.toDos, t)
	return nil
}

func (r *ToDoRepositoryInMemory) Delete(id string) error {
	for i, u := range r.toDos {
		if u.ID == id {
			r.toDos = append(r.toDos[:i], r.toDos[i+1:]...)
			return nil
		}
	}
	return errors.New("TO DO not found")
}
