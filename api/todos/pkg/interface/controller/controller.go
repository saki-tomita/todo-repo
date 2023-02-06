package controller

//ルーターからルーティングを受け取り、ユースケースのInteractorに渡す
/*
import (
	"todos/pkg/interface/database"
	"todos/pkg/usecase"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(DBConnector database.DBConnector) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				DBConnector: DBConnector,
			},
		},
	}
}

func (controller *TaskController) ShowTasks() {
	userId := 1 //httpリクエストで渡すemailから取得するユーザーidとする
	_, err := controller.Interactor.Tasks(userId)
	if err != nil {
		//500エラーを返す処理
		return
	}
	return
}

func Testfunc() {}
*/
