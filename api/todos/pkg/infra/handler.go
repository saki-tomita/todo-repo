package infra

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"todos/pkg/domain"
	"todos/pkg/usecase/task"
)

func GetAllTask(Interactor task.TaskRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetAllTask")
		log.Println(r.Host + r.RequestURI)
		var tasks []*domain.Task
		var err error

		email := GetParam(r, "email")
		log.Println(email)

		//サービス層を呼びリストを取得
		tasks, err = Interactor.GetAllTask(email)
		log.Println("hanler")
		log.Println(tasks)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
			return
		}

		if tasks == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error reading tasks"))
			return
		}

		//JSON形式変換
		var toJ []*domain.Jtask
		for _, t := range tasks {
			toJ = append(toJ, &domain.Jtask{
				ID:       t.ID,
				Name:     t.Name,
				Status:   t.Status,
				Rank:     t.Rank,
				Deadline: t.Deadline,
				Label:    t.Label,
				Memo:     t.Memo,
				UserID:   t.UserID,
				DelFlg:   t.DelFlg,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
		}
	})
}

func GetTask(Interactor task.TaskRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//パラメータ取得
		query := MakeQuery(r)
		//サービス層呼びだし
		task, err := Interactor.GetTask(query)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
			return
		}

		var toJ []*domain.Jtask
		for _, t := range task {
			toJ = append(toJ, &domain.Jtask{
				ID:       t.ID,
				Name:     t.Name,
				Status:   t.Status,
				Rank:     t.Rank,
				Deadline: t.Deadline,
				Label:    t.Label,
				Memo:     t.Memo,
				UserID:   t.UserID,
				DelFlg:   t.DelFlg,
			})
		}

		//JSON変換
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading tasks"))
		}
	})
}

func createTask(Interactor task.TaskRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("handler createTask")
		log.Println(r.Body)
		errMessage := "Error adding task"
		var input domain.Task //Jsonの方定義でつまりそう
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		id, err := Interactor.CreateTask(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		toJ := &domain.Jtask{
			ID:       id,
			Name:     input.Name,
			Status:   input.Status,
			Rank:     input.Rank,
			Deadline: input.Deadline,
			Label:    input.Label,
			Memo:     input.Memo,
			UserID:   input.UserID,
			DelFlg:   input.DelFlg,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
	})
}

func updateTask(Interactor task.TaskRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error reading tasks"
		//パラメータ取得
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		//更新
		var input domain.Task
		err2 := json.NewDecoder(r.Body).Decode(&input)
		if err2 != nil {
			log.Println(err2.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		input.ID = id
		err = Interactor.UpdateTask(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		toJ := &domain.Jtask{
			ID:       id,
			Name:     input.Name,
			Status:   input.Status,
			Rank:     input.Rank,
			Deadline: input.Deadline,
			Label:    input.Label,
			Memo:     input.Memo,
			UserID:   input.UserID,
			DelFlg:   input.DelFlg,
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
	})
}

func deleteTask(Interactor task.TaskRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errMessage := "Error deleting tasks"
		//パラメータ取得
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		//削除（del_id更新）
		err = Interactor.DeleteTask(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errMessage))
			return
		}
		w.WriteHeader(http.StatusOK)

	})
}

// ********** パラメータ取り出し系_後々ファイル分けたい *************//
func MakeQuery(r *http.Request) string {
	var UriFull string
	UriFull = "http://" + r.Host + r.RequestURI
	u, _ := url.Parse(UriFull) //URL自体を取得する方法がわからない
	//u, _ := url.Parse("https://xxx.com?a=AAA&b=BBB&c=CCC&d=DDD")
	vars := u.Query()
	var query string
	var q string
	for key, v := range vars {
		if key == "email" {
			q = "t_user." + key + "=" + "'" + v[0] + "'"
		} else {
			q = "t_task." + key + "=" + "'" + v[0] + "'"
		}
		if query != "" {
			query += " and "
		}
		query += q
	}
	return query
}

func GetParam(r *http.Request, pKey string) string {
	query := r.URL.Query()
	return query.Get(pKey)
}

//**********************************************************//

func TaskHandlers(r *mux.Router, n negroni.Negroni, interactor task.TaskRepository) {

	//r.HandleFunc("/task", Cors(GetAllTask, interactor)).Methods("GET", "OPTIONS")

	r.Handle("/task", n.With(
		negroni.Wrap(GetAllTask(interactor)),
	)).Methods("GET", "OPTIONS").Queries("email", "{email}").Name("GetAllTask")

	/*
		r.Handle("/task", n.With(
			negroni.Wrap(Cors(GetAllTask, interactor)),
		)).Methods("GET", "OPTIONS").Queries("email", "{email}").Name("GetAllTask")

	*/

	r.Handle("/task/{id}", n.With(
		negroni.Wrap(GetTask(interactor)),
	)).Methods("GET", "OPTIONS").Queries("name", "{name}").Queries("status", "{status}").Name("GetTask")

	r.Handle("/task", n.With(
		negroni.Wrap(createTask(interactor)),
	)).Methods("POST", "OPTIONS").Name("createTask")

	r.Handle("/task/{id}", n.With(
		negroni.Wrap(updateTask(interactor)),
	)).Methods("PUT", "OPTIONS").Name("updateTask")

	r.Handle("/task/del/{id}", n.With(
		negroni.Wrap(deleteTask(interactor)),
	)).Methods("PUT", "OPTIONS").Name("deleteTask")

	//疎通確認
	r.HandleFunc("/", hello_world)
}

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
