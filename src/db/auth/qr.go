package auth

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/db/auth/internal/qrmanager"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/internal/store"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
)

func GenerateQr(qr model.QrAuth) (string, error) {
	qrByte, err := qrmanager.GenerateQr(qr.QrCode)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	if err = store.InsertQr(db, qr); err != nil {
		return "", err
	}

	qrBase64 := base64.StdEncoding.EncodeToString(qrByte)
	return qrBase64, nil
}

func FetchQr(userId string) (model.QrAuth, error) {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	qrAuth, err := store.FetchQrByUserId(db, userId)
	if err != nil {
		return qrAuth, err
	}
	return qrAuth, nil
}
