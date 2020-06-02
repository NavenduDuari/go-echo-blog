package main

import (
	"database/sql"
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
)

func InstallPsqlUUIDExtension(db *sql.DB) error {
	_, err := db.Query("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		log.Printf("Not able to create extension uuid-ossp", err)
	}
	return err
}

func SetTimeZone(db *sql.DB) error {
	setTimeZone := "SET timezone = 'Asia/Kolkata';"
	_, err := db.Exec(setTimeZone)
	if err != nil {
		return err
	}
	return nil
}

func CreatePsqlTableUser(db *sql.DB) (*sql.Rows, error) {

	createTable := "CREATE TABLE IF NOT EXISTS \"" + constant.PostgressTableUser + "\" (" +
		`id UUID NOT NULL DEFAULT uuid_generate_v1(),
		email TEXT UNIQUE,
		phone TEXT UNIQUE,
		password TEXT NOT NULL
		);`
	rows, err := db.Query(createTable)
	if err != nil {
		return nil, err
	}
	return rows, err
}
