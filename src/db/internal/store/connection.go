package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	_ "github.com/lib/pq"
)

var db *sql.DB
var isDbInitialized = false

func InitPostgressDB() {
	pgsqlUser := os.Getenv("YO_PGSQL_BLOG_USER")
	pgsqlPassword := os.Getenv("YO_PGSQL_BLOG_PASSWORD")
	pgsqlHost := os.Getenv("YO_PGSQL_BLOG_HOST")
	pgsqlPort := os.Getenv("YO_PGSQL_BLOG_PORT")
	pgsqlDbname := constant.PostgressBlogDBName

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgsqlHost, pgsqlPort, pgsqlUser, pgsqlPassword, pgsqlDbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil || db == nil {
		// yolog.ErrorAndExit("Unable to connect "+constant.PostgressBlogDBName, err)
	}

	err = db.Ping()
	if err != nil {
		// yolog.ErrorAndExit("Unable to connect "+yoconst.PostgressYoblogDBName, err)

	}
	// yolog.Info("Database connected : " + yoconst.PostgressYoblogDBName)
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
