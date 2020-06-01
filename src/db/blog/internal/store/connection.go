package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	_ "github.com/lib/pq"
)

var db *sql.DB
var isDbInitialized = false

func InitPostgressDB() {
	pgsqlUser := os.Getenv("PGSQL_BLOG_USER")
	pgsqlPassword := os.Getenv("PGSQL_BLOG_PASSWORD")
	pgsqlHost := os.Getenv("PGSQL_BLOG_HOST")
	pgsqlPort := os.Getenv("PGSQL_BLOG_PORT")
	pgsqlDbname := constant.PostgressBlogDBName

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgsqlHost, pgsqlPort, pgsqlUser, pgsqlPassword, pgsqlDbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil || db == nil {
		log.Fatal("Unable to connect "+constant.PostgressBlogDBName, err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect "+constant.PostgressBlogDBName, err)

	}
	fmt.Println("Database connected : " + constant.PostgressBlogDBName)
	isDbInitialized = true
}

func GetPostgressDB() (*sql.DB, error) {
	if !isDbInitialized || db == nil {
		return nil, errors.New("Database is not initialized : " + constant.PostgressBlogDBName)
	}
	return db, nil
}

func ClosePostgressDB() {
	db.Close()
}
