package service

import (
	"HGMchat/internal/dao"
	"HGMchat/internal/model"
	"HGMchat/pkg"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) error {
	if dao.IsUsernameExists(username) {
		return errors.New("用户名已存在")
	}

	HashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: string(HashPwd),
	}

	return dao.Create(user)

}

func Login(username, password string) (string, error) {
	user, err := dao.GetByUsername(username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	token, err := pkg.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
