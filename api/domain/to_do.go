package domain

type Status string

const (
	Completed Status = "Completed"
	Pending   Status = "Pending"
)

type ToDo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type ToDoRepository interface {
	GetAll() ([]ToDo, error)
	Add(t ToDo) error
	Delete(id string) error
}
