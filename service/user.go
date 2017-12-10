package service

import "github.com/bitschain/panicgo/model"

func (s *PanicService) CreateNewUser(data model.AccountInfo) (model.User, error) {
	var user model.User
	user.Name = data.Name
	user.Email = data.Email
	user.Password = user.EncryptPassword(data.Password, user.Salt())
	user.Role = model.UserRoleNormal
	user.Status = model.UserStatusInActive

	newUser, err := s.Model.CreateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
