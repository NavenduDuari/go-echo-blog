package auth

import "github.com/NavenduDuari/go-echo-blog/src/db/auth/internal/store"

func InitStore() {
	store.InitPostgressDB()
}

func CloseStore() {
	store.ClosePostgressDB()
}
