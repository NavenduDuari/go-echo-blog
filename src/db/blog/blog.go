package blog

import (
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/db/blog/internal/store"
	"github.com/NavenduDuari/go-echo-blog/src/db/blog/model"
)

func InsertBlog(newBlog *model.Blog) error {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	err = store.InsertBlog(db, newBlog)
	if err != nil {
		return err
	}
	return nil
}

func FetchBlogs() ([]model.Blog, error) {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	ads, err := store.FetchBlogs(db)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func DeleteBlog(id string) error {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	err = store.DeleteBlog(db, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBlog(id string, updatedBlog model.Blog) error {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	err = store.UpdateBlog(db, id, updatedBlog)
	if err != nil {
		return err
	}
	return nil
}
