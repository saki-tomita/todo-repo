package adapter

import (
	_ "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"todos/pkg/infra"
	"todos/pkg/usecase/task"
)

func InitRouter() {
	taskConn := infra.NewDbConnector()
	defer taskConn.DbPool.Close()
	taskInteractor := task.NewInteractor(taskConn)

	router := mux.NewRouter()

	n := negroni.New(
		//negroni.HandlerFunc(Cors),
		negroni.NewLogger(),
	)

	TaskHandlers(router, *n, taskInteractor)
	http.Handle("/", router)
	log.Println("Lintening server '0.0.0.0:8081'")

	server := http.Server{
		Addr: "0.0.0.0:8081",
		//Handler:      handler,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
