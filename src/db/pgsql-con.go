package yoblog

import "github.com/NavenduDuari/go-echo-blog/src/db/internal/store"

func InitStore() {
	store.InitPostgressDB()
}

func CloseStore() {
	store.ClosePostgressDB()
}
