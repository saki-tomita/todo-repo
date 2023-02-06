//APIの実処理を記述
//repository層のインターフェースを実装

package database

/*
import (
	"todos/pkg/domain"
)


type TaskRepository struct {
	DBConnector
}

func (repo *TaskRepository) GetTasks(userId int) (tasks domain.Tasks, err error) {
	row, err := repo.Query(
		"SELECT id, name, status, rank, deadline, label, memo, user_id, del_flg"+
			"FROM t_user WHERE user_id = $1", userId)
	defer row.Close()
	if err != nil {
		return
	}
	for row.Next() {
		var id int
		var name string
		var status int
		var rank int
		var deadline string
		var label int
		var memo string
		var usId int
		var dFlg int
		if err := row.Scan(&id, &name, &status, &rank, &deadline, &label, &memo, &usId, &dFlg); err != nil {
			return
		}
		task := domain.Task{
			ID:       id,
			Name:     name,
			Status:   status,
			Rank:     rank,
			Deadline: deadline,
			Label:    label,
			Memo:     memo,
			UserID:   usId,
			DqlFlg:   dFlg,
		}
		tasks = append(tasks, task)
	}
	return
}

//クエリ試し（あとでAPIにうつす）
var (
	email string
)
//ここからおかしい
infra.err = infra.dbPool.QueryRow("SELECT email FROM t_user WHERE id = $1", 1).Scan(&email)
if infra.err != nil {
return nil
}
fmt.Println(email)

return email

*/
