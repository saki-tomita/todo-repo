package infra

//DB接続
//interface/go_handlerのインターフェース定義を実装
//外部から提供されるsql.Rowsなどをラップ

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"time"
	"todos/pkg/domain"
	//"todos/pkg/interface/database"

	_ "cloud.google.com/go/cloudsqlconn"
	_ "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v5"
)

type DBConnector struct {
	dbPool *sql.DB
}

func (conn *DBConnector) Get(q string) ([]*domain.Task, error) {
	query := "select t_task.* from t_task inner join t_user on t_task.user_id = t_user.id where " + q
	log.Println(query)
	stmt, err := conn.dbPool.Prepare(query)
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	rows, err := stmt.Query()
	log.Println("rows")
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		task := &domain.Task{}
		log.Println("task")
		log.Println(task)
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.Rank, &task.Deadline, &task.Label, &task.Memo, &task.UserID, &task.DelFlg, &task.CreatedAt, &task.CreatedUser, &task.UpdatedAt, &task.UpdatedUser)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	log.Println("db_connector")
	log.Println(tasks)
	return tasks, nil
}

func (conn *DBConnector) GetAll(email string) ([]*domain.Task, error) {
	query := "select t_task.* from t_task inner join t_user on t_task.user_id = t_user.id where t_task.del_flg = 0 and t_user.email = $1"
	stmt, err := conn.dbPool.Prepare(query)
	if err != nil {
		return nil, err
	}
	var tasks []*domain.Task
	rows, err := stmt.Query(email)
	log.Println("rows")
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	//var task domain.Task
	for rows.Next() {
		task := &domain.Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.Rank, &task.Deadline, &task.Label, &task.Memo, &task.UserID, &task.DelFlg, &task.CreatedAt, &task.CreatedUser, &task.UpdatedAt, &task.UpdatedUser)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	log.Println("db_connector")
	log.Println(tasks)
	return tasks, nil
}

func (conn *DBConnector) Create(task *domain.Task) (int, error) {
	log.Println("db_connector Create")
	query := "insert into t_task (name, status, rank, deadline, label, memo, user_id, del_flg, created_at, created_user, updated_at, updated_user) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	log.Println(query)
	stmt, err := conn.dbPool.Prepare(query)
	if err != nil {
		return task.ID, err
	}
	_, err = stmt.Exec(task.Name, task.Status, task.Rank, task.Deadline, task.Label, task.Memo, task.UserID, task.DelFlg, time.Now(), "postgres", time.Now(), "postgres")
	if err != nil {
		return task.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return task.ID, err
	}
	return task.ID, nil
}

func (conn *DBConnector) Update(task *domain.Task) error {
	log.Println("db_connector Update")
	query := "update t_task set name = $1, status = $2, rank = $3, deadline = $4, label = $5, memo = $6, user_id = $7, del_flg = $8, updated_at = $9, updated_user = $10 where id = $11"
	log.Println(query)
	stmt, err := conn.dbPool.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(task.Name, task.Status, task.Rank, task.Deadline, task.Label, task.Memo, task.UserID, task.DelFlg, time.Now(), "postgres", task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (conn *DBConnector) Delete(id int) error {
	query := "update t_task set del_flg = 1 where id = $1;"
	stmt, err := conn.dbPool.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func NewDbConnector() *DBConnector {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		dbUser                 = os.Getenv("DB_USER")                   // e.g. 'my-db-user'
		dbIAMUser              = os.Getenv("DB_IAM_USER")               // e.g. 'sa-name@project-id.iam'
		dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
		dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		usePrivate             = os.Getenv("PRIVATE_IP")
	)
	if dbUser == "" && dbIAMUser == "" {
		log.Fatal("Warning: One of DB_USER or DB_IAM_USER must be defined")
	}
	dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	var opts []cloudsqlconn.Option
	if dbIAMUser != "" {
		opts = append(opts, cloudsqlconn.WithIAMAuthN())
	}
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
	}
	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		panic(err)
	}
	// Use the Cloud SQL connector to handle connecting to the instance.
	// This approach does *NOT* require the Cloud SQL proxy.
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	var dbPool *sql.DB
	dbPool, err = sql.Open("pgx", dbURI)
	if err != nil {
		panic(err)
	}

	//DBConnector := new(DBConnector)
	//DBConnector.dbPool = dbPool
	return &DBConnector{
		dbPool: dbPool,
	}
}
