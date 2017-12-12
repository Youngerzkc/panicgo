package service

import (
	"github.com/bitschain/panicgo/errors"
	"github.com/bitschain/panicgo/model"
	"github.com/go-sql-driver/mysql"
	"log"
	"github.com/gin-gonic/gin"
)

func (s *PanicService) RegisterUser(data model.UserRegister) (model.User, *errors.PanicError) {
	var user model.User
	user.Name = data.Name
	user.Email = data.Email
	user.Password = user.EncryptPassword(data.Password, user.Salt())
	user.Role = model.UserRoleNormal
	user.Status = model.UserStatusInActive

	newUser, err := s.Model.CreateUser(user)

	if sqlErr, ok := err.(*mysql.MySQLError); ok {
		if sqlErr.Number == 1062 {

			return newUser, errors.Wrap(err, 1000, "邮箱已注册")
		}
	}

	if err != nil {
		return newUser, errors.Wrap(err, 1001, "注册用户失败")
	}
	return newUser, nil
}

func (s *PanicService) Login(data model.UserLogin) {

}
