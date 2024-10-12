package usecase

import "api/domain"

type ToDoUsecase struct {
	repo domain.ToDoRepository
}

func NewToDoUsecase(repo domain.ToDoRepository) *ToDoUsecase {
	return &ToDoUsecase{repo: repo}
}

func (tu *ToDoUsecase) GetTodos() ([]domain.ToDo, error) {
	return tu.repo.GetAll()
}

func (tu *ToDoUsecase) CreateToDo(t domain.ToDo) error {
	return tu.repo.Add(t)
}

func (tu *ToDoUsecase) DeleteToDo(id string) error {
	return tu.repo.Delete(id)
}
