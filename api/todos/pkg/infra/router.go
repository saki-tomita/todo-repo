package infra

import (
	_ "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todos/pkg/usecase/task"
)

func InitRouter() {
	taskConn := NewDbConnector()
	defer taskConn.dbPool.Close()
	taskInteractor := task.NewInteractor(taskConn)

	router := mux.NewRouter()

	//router.HandleFunc("/task", Cors(TaskHandlers)).Methods("GET", "OPTIONS")

	n := negroni.New(
		//negroni.HandlerFunc(Cors),
		negroni.NewLogger(),
	)

	TaskHandlers(router, *n, taskInteractor)
	http.Handle("/", router)
	log.Println("Lintening server '0.0.0.0:8081'")

	// CORS レスポンスヘッダーの追加
	/*
		c := cors.Default()
		handler := c.Handler(router)
	*/
	//handler := context.ClearHandler(http.DefaultServeMux)
	server := http.Server{
		Addr: "0.0.0.0:8081",
		//Handler:      handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

/*
func getTaskhandle(w http.ResponseWriter, r *http.Request) {
	// DB(CloudSQLに接続)
	dbConn := NewDbConnector()
	//httpサーバー
	taskController := controller.NewTaskController(dbConn)
	taskController.ShowTasks()
}
*/
