// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5"
)

var dbPool *sql.DB
var err error // *sql.DBだけでなく、errもpackageグローバルに定義する

func main() {

	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)

	//DBクエリ実行確認
	fmt.Println("Calling infra...")
	//infra.ConnectWithConnector()

	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}

	var (
		dbUser = mustGetenv("DB_USER") // e.g. 'my-db-user'
		dbPwd  = mustGetenv("DB_PASS") // e.g. 'my-db-password'
		//dbTCPHost = mustGetenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
		dbTCPHost = "127.0.0.1"
		//instanceConnectionName = "/cloudsql/ca-saki-tomita-test:asia-northeast1:postgres-todo-instance"
		//dbPort = mustGetenv("DB_PORT") // e.g. '5432'
		dbName = mustGetenv("DB_NAME") // e.g. 'my-database'
	)

	if dbUser == "" {
		log.Fatal("Warning: One of DB_USER must be defined")
	}

	dbURI := fmt.Sprintf("host=%s user=%s password=%s database=%s",
		dbTCPHost, dbUser, dbPwd, dbName)
	//dbURI := "postgres: //postgres:oremeka9@127.0.0.1:5432/tododev?sslmode=disable"
	fmt.Println("dbURI:")
	fmt.Println(dbURI)

	if dbRootCert, ok := os.LookupEnv("DB_ROOT_CERT"); ok { // e.g., '/path/to/my/server-ca.pem'
		var (
			dbCert = mustGetenv("DB_CERT") // e.g. '/path/to/my/client-cert.pem'
			dbKey  = mustGetenv("DB_KEY")  // e.g. '/path/to/my/client-key.pem'
		)
		dbURI += fmt.Sprintf(" sslmode=require sslrootcert=%s sslcert=%s sslkey=%s",
			dbRootCert, dbCert, dbKey)
	}

	dbPool, err := sql.Open("pgx", dbURI)
	err = dbPool.Ping()

	fmt.Println("dbPool:")
	fmt.Println(dbPool)

	//クエリ試し（あとでAPIにうつす）
	var (
		email string
	)
	//ここからおかしい
	err = dbPool.QueryRow("SELECT email FROM users WHERE id = $1", 1).Scan(&email)
	if err != nil {
		return
	}

	fmt.Println(email)

}
