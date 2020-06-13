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

func CreatePsqlTableOtpAuth(db *sql.DB) (*sql.Rows, error) {
	createTable := "CREATE TABLE IF NOT EXISTS \"" + constant.PostgressTableOtpAuth + "\" (" +
		`id UUID NOT NULL DEFAULT uuid_generate_v1(),
		userid TEXT UNIQUE,
		otp TEXT UNIQUE,
		expires_at TEXT NOT NULL
		);`
	rows, err := db.Query(createTable)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func CreatePsqlTableQrAuth(db *sql.DB) (*sql.Rows, error) {
	createTable := "CREATE TABLE IF NOT EXISTS \"" + constant.PostgressTableQrAuth + "\" (" +
		`id UUID NOT NULL DEFAULT uuid_generate_v1(),
		userid TEXT UNIQUE,
		qr_code TEXT UNIQUE,
		expires_at TEXT NOT NULL
		);`
	rows, err := db.Query(createTable)
	if err != nil {
		return nil, err
	}
	return rows, err
}
