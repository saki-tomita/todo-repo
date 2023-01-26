//DB接続

package infra

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"log"
	"net"
	"os"

	_ "cloud.google.com/go/cloudsqlconn"
	_ "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v5"
)

var dbPool *sql.DB

var err error // *sql.DBだけでなく、errもpackageグローバルに定義する

func ConnectWithConnector() (*sql.DB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep secrets safe.

	var (
		// Either a DB_USER or a DB_IAM_USER should be defined. If both are
		// defined, DB_IAM_USER takes precedence.
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
		return nil, err
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
		return nil, err
	}
	// Use the Cloud SQL connector to handle connecting to the instance.
	// This approach does *NOT* require the Cloud SQL proxy.
	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceConnectionName)
	}
	dbURI := stdlib.RegisterConnConfig(config)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	//クエリ試し（あとでAPIにうつす）
	var (
		email string
	)
	//ここからおかしい
	err = dbPool.QueryRow("SELECT email FROM t_user WHERE id = $1", 1).Scan(&email)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	fmt.Println(email)

	return dbPool, nil
}
