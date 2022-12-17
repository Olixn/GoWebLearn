package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/Olixn/GoWebLearn/models"
)

const secret = "golang111"

// 把每一步数据库操作封装成函数
// 等待logic层根据业务需求调用

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	} else if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对用户密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `INSERT INTO user(user_id, username, password) VALUES (?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `SELECT user_id,username,password FROM user WHERE username = ?`
	if err = db.Get(user, sqlStr, user.Username); err == sql.ErrNoRows {
		return ErrorUserNotExist
	} else if err != nil {
		// 数据库查询出错
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func GetUserByID(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `SELECT user_id,username FROM user WHERE user_id = ?`
	err = db.Get(user, sqlStr, id)
	return
}
