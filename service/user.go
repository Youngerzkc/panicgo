package service

import (
	"github.com/bitschain/panicgo/errors"
	"github.com/bitschain/panicgo/model"
	"github.com/go-sql-driver/mysql"
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

func (s *PanicService) Login(data model.UserLogin) error {
		 var user model.User
		 	
		 user.Email=data.Email
		 var err error
 		if user,err =s.Model.QueryUser(user);err!=nil{
		
			return errors.Wrap(err,1002,"该用户未注册")
		 }
		 passwd:=user.EncryptPassword(data.Password, user.Salt()) 
		 if passwd != user.Password{
			return errors.Wrap(err, 1003, "密码错误,请重新输入密码")
		 }
		 return nil
}
