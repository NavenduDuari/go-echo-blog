package main

import (
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	"github.com/NavenduDuari/go-echo-blog/src/db/blog/internal/store"
)

func main() {
	store.InitPostgressDB()

	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}

	err = InstallPsqlUUIDExtension(db)
	if err != nil {
		log.Fatal("Unable to install UUID extension on "+constant.PostgressBlogDBName, err)
	}

	err = SetTimeZone(db)
	if err != nil {
		log.Fatal("Unable to set time zone on "+constant.PostgressBlogDBName, err)
	}

	_, err = CreatePsqlTableBlog(db)
	if err != nil {
		log.Fatal("Unable to create table: "+constant.PostgressBlogDBName, err)
	}
}
