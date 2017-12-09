package service

import "github.com/bitschain/panicgo/model"

type AccountInfo struct {
	Name     string `json:"name" valid:"runelength(4|20)"`
	Email    string `json:"email" valid:"email,runelength(5|50)"`
	Password string `json:"password" valid:"runelength(6|20)"`
}

func (s *PanicService) CreateNewUser(data AccountInfo) (model.User, error) {
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
