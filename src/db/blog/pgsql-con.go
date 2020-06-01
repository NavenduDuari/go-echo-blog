package blog

import "github.com/NavenduDuari/go-echo-blog/src/db/blog/internal/store"

func InitStore() {
	store.InitPostgressDB()
}

func CloseStore() {
	store.ClosePostgressDB()
}
