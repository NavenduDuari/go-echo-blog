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

func CreatePsqlTableBlog(db *sql.DB) (*sql.Rows, error) {

	createType := `CREATE TYPE "body_component" AS (
		title TEXT,
		content TEXT,
		images TEXT[]
		);`
	rows, err := db.Query(createType)
	if err != nil {
		return nil, err
	}
	createTable := "CREATE TABLE IF NOT EXISTS \"" + constant.PostgressTableBlog + "\" (" +
		`id UUID NOT NULL DEFAULT uuid_generate_v1(),
		title TEXT NOT NULL,
		body "body_component",
		author TEXT NOT NULL,
		timestamp TIMESTAMPTZ,
		categories TEXT[],
		tags TEXT[]
		);`
	rows, err = db.Query(createTable)
	if err != nil {
		return nil, err
	}
	return rows, err
}
