package store

import (
	"database/sql"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
)

//TODO delete qr automatically after certain time
func InsertQr(db *sql.DB, newQr model.QrAuth) error {
	_, err := db.Exec("INSERT INTO \""+constant.PostgressTableQrAuth+
		"\" (userid, qr_code, expires_at) VALUES($1, $2, $3) ON CONFLICT(userid) DO UPDATE SET qr_code = $2, expires_at = $3",
		newQr.UserId,
		newQr.QrCode,
		newQr.ExpiresAt)

	if err != nil {
		return err
	}
	return nil
}

func FetchQrByUserId(db *sql.DB, userId string) (model.QrAuth, error) {
	var qr model.QrAuth

	rows, err := db.Query("Select * from \""+constant.PostgressTableQrAuth+"\" WHERE userid = $1", userId)
	if err != nil || rows == nil {
		return qr, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&qr.Id, &qr.UserId, &qr.QrCode, &qr.ExpiresAt)
		if err != nil {
			return qr, err
		}
	}
	return qr, nil
}

// func FetchUserByEmail(db *sql.DB, email string) (model.User, error) {
// 	var user model.User

// 	rows, err := db.Query("Select * from \""+constant.PostgressTableOtpAuth+"\" WHERE email = $1", email)
// 	if err != nil || rows == nil {
// 		return user, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&user.Id, &user.Email, &user.Phone, &user.Password)
// 		if err != nil {
// 			return user, err
// 		}
// 	}
// 	return user, nil
// }

// func DeleteUser(db *sql.DB, id string) error {
// 	_, err := db.Exec("DELETE from "+constant.PostgressTableOtpAuth+" WHERE id = $1", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
