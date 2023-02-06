package task

import (
	"log"
	"todos/pkg/domain"
)

//Infra層のファンクション呼び出しを表すRepository
//Repositoryをユースケース単位に内包するTaskRepository
//interfaces/databaseからのInput Portの役割、およびinterfaces/controllersへのGatewayの役割

type Interactor struct {
	repo Repository
}

func NewInteractor(r Repository) *Interactor {
	return &Interactor{
		repo: r,
	}
}

func (i Interactor) GetTask(query string) ([]*domain.Task, error) {
	task, err := i.repo.Get(query)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (i Interactor) GetAllTask(email string) ([]*domain.Task, error) {
	//service→repositoryのList()→infraのList()を呼び出す
	//usecase層からinfra層を呼び出すため
	tasks, err := i.repo.GetAll(email)
	if err != nil {
		return nil, err
	}
	log.Println("interactor")
	log.Println(tasks)
	return tasks, nil
}

func (i Interactor) CreateTask(t *domain.Task) (int, error) {
	//Task構造体を設定値で定義
	task, err := domain.NewTask(*t)
	if err != nil {
		return task.ID, err
	}
	log.Println("Interactor CreateTask")
	return i.repo.Create(task)
}

func (i Interactor) UpdateTask(t *domain.Task) error {
	err := i.repo.Update(t)
	if err != nil {
		return err
	}
	return nil
}

func (i Interactor) DeleteTask(id int) error {
	err := i.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
