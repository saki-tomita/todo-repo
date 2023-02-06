package task

import "todos/pkg/domain"

//Infra層のファンクション呼び出しを表すRepository
//Repositoryをユースケース単位に内包するTaskRepository

type Repository interface {
	Get(query string) ([]*domain.Task, error)
	GetAll(email string) ([]*domain.Task, error)
	Create(task *domain.Task) (int, error)
	Update(task *domain.Task) error
	Delete(id int) error
}

type TaskRepository interface {
	GetTask(query string) ([]*domain.Task, error)
	GetAllTask(email string) ([]*domain.Task, error)
	CreateTask(task *domain.Task) (int, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id int) error
}
