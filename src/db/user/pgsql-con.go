package user

import "github.com/NavenduDuari/go-echo-blog/src/db/user/internal/store"

func InitStore() {
	store.InitPostgressDB()
}

func CloseStore() {
	store.ClosePostgressDB()
}
