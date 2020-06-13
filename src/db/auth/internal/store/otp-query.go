package store

import (
	"database/sql"

	"github.com/NavenduDuari/go-echo-blog/src/common/constant"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
)

//TODO delete otp automatically after certain time
func InsertOtp(db *sql.DB, newOtp model.OtpAuth) error {
	_, err := db.Exec("INSERT INTO \""+constant.PostgressTableOtpAuth+
		"\" (userid, otp, expires_at) VALUES($1, $2, $3) ON CONFLICT(userid) DO UPDATE SET otp = $2, expires_at = $3",
		newOtp.UserId,
		newOtp.OTP,
		newOtp.ExpiresAt)

	if err != nil {
		return err
	}
	return nil
}

func FetchOtpByUserId(db *sql.DB, userId string) (model.OtpAuth, error) {
	var otp model.OtpAuth

	rows, err := db.Query("Select * from \""+constant.PostgressTableOtpAuth+"\" WHERE userid = $1", userId)
	if err != nil || rows == nil {
		return otp, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&otp.Id, &otp.UserId, &otp.OTP, &otp.ExpiresAt)
		if err != nil {
			return otp, err
		}
	}
	return otp, nil
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
