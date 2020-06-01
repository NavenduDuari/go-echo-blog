package store

import (
	"database/sql"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	"github.com/NavenduDuari/go-echo-blog/src/db/user/model"
)

func InsertUser(db *sql.DB, newUser *model.User) error {
	_, err := db.Exec("INSERT INTO "+constant.PostgressTableUser+
		" (phone, email, password) VALUES($1, $2, $3)",
		newUser.Phone,
		newUser.Email,
		newUser.Password)

	if err != nil {
		return err
	}
	return nil
}

func FetchUserById(db *sql.DB, id string) (model.User, error) {
	var user model.User

	rows, err := db.Query("Select * from "+constant.PostgressTableUser+" WHERE id = $1", id)
	if err != nil || rows == nil {
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Phone, &user.Password)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE from "+constant.PostgressTableUser+" WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
