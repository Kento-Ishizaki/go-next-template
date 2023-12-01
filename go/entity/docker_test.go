package entity_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	user     = "root"
	password = "password"
	dbName   = "test_db"
	port     = 3307
	dialect  = "mysql"
)

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	// pool.MaxWait = time.Second * 20
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + password,
			"MYSQL_DATABASE=" + dbName,
		},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: "0.0.0.0", HostPort: "3307"},
			},
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions,
		func(config *docker.HostConfig) {
			config.AutoRemove = false
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 10)

		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, port, dbName)

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

func setupTestDB(db *sql.DB) {
	mgrsPath, _ := filepath.Abs("../db/migrations")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+mgrsPath,
		dialect,
		driver,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlDir: %v, error: %v\n", mgrsPath, err)
		os.Exit(1)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("no change made by migration scripts")
		} else {
			log.Fatal(err)
		}
	}
}

func TestMain(m *testing.M) {
	resource, pool := createContainer()
	defer closeContainer(resource, pool)

	db := connectDB(resource, pool)
	defer db.Close()

	// uses pool to try to connect to Docker
	err := pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// テスト用DBをセットアップ
	setupTestDB(db)
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestSomething(t *testing.T) {
	expect := 1
	actual := 1
	assert.Equal(t, expect, actual)
}
