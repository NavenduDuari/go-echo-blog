package user

import (
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/db/user/internal/store"
	"github.com/NavenduDuari/go-echo-blog/src/db/user/model"
)

func InsertUser(newUser *model.User) error {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	err = store.InsertUser(db, newUser)
	if err != nil {
		return err
	}
	return nil
}

func FetchUserByUserId(userId string) (model.User, error) {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	user, err := store.FetchUserByPhone(db, userId)
	if err != nil {
		user, err = store.FetchUserByEmail(db, userId)
		if err != nil {
			return model.User{}, err
		}
	}
	return user, nil
}

func DeleteUser(id string) error {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	err = store.DeleteUser(db, id)
	if err != nil {
		return err
	}
	return nil
}
