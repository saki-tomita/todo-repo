//APIでやりとりするドメインモデル
//構造体、カラム定義など

package domain

type User struct {
	ID    int
	Email string
}

type Task struct {
	ID          int
	Name        string
	Status      int
	Rank        int
	Deadline    string
	Label       int
	Memo        string
	UserID      int
	DelFlg      int
	CreatedAt   string
	CreatedUser string
	UpdatedAt   string
	UpdatedUser string
}

type Jtask struct {
	ID       int `json:"ID,string"`
	Name     string
	Status   int `json:"Status,string"`
	Rank     int `json:"Rank,string"`
	Deadline string
	Label    int `json:"Label,string"`
	Memo     string
	UserID   int `json:"UserID,string"`
	DelFlg   int `json:"DelFlg,string"`
}

func NewTask(task Task) (*Task, error) {
	t := &Task{
		ID:          task.ID,
		Name:        task.Name,
		Status:      task.Status,
		Rank:        task.Rank,
		Deadline:    task.Deadline,
		Label:       task.Label,
		Memo:        task.Memo,
		UserID:      task.UserID,
		DelFlg:      task.DelFlg,
		CreatedAt:   task.CreatedAt,
		CreatedUser: task.CreatedUser,
		UpdatedAt:   task.UpdatedAt,
		UpdatedUser: task.UpdatedUser,
	}
	return t, nil
}
