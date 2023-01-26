// main.goから呼びだされ、各層の橋渡しを担当
package http

import (
	"github.com/gorilla/mux"
	infra "todos/pkg/infra"
)

func initRouter() {
	// DB(CloudSQLに接続)
	dbPool, _ := infra.connectWithConnector()

	//SQL試し（後々インフラ層にうつす）

	r := mux.NewRouter()
	r.HandleFunc("/getTask", infra.getAllTaskList)
}
