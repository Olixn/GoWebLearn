package logic

import (
	"github.com/Olixn/GoWebLearn/dao/mysql"
	"github.com/Olixn/GoWebLearn/models"
	"github.com/Olixn/GoWebLearn/pkg/jwt"
	"github.com/Olixn/GoWebLearn/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 传的是指针，所以在mysql包查询成功并赋值，logic包中的user也变了
	if err = mysql.Login(user); err != nil {
		return nil, err
	}

	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
