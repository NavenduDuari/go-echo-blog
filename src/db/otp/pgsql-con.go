package otp

import "github.com/NavenduDuari/go-echo-blog/src/db/otp/internal/store"

func InitStore() {
	store.InitPostgressDB()
}

func CloseStore() {
	store.ClosePostgressDB()
}
