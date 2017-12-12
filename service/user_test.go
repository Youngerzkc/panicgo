package service

import (
	"github.com/bitschain/panicgo/model"
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicService_CreateNewUser(t *testing.T) {
	s := GetTestService()

	fakeUser, _ := faker.New("en")
	data := model.UserRegister{Name: fakeUser.Name(), Email: fakeUser.Email(), Password: "123456"}

	user, err := s.RegisterUser(data)
	assert.Nil(t, err)
	assert.Equal(t, data.Name, user.Name)
	assert.Equal(t, data.Email, user.Email)
}
