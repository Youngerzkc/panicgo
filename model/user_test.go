package model

import (
	"fmt"
	"testing"

	"github.com/bitschain/panicgo/database"
	"github.com/stretchr/testify/assert"
)

func TestSalt(t *testing.T) {
	var user User
	fmt.Println(user.Salt())
	assert.Equal(t, len(user.Salt()), 10)
}

func TestEncryptPassword(t *testing.T) {
	var user User
	hashPwd := user.EncryptPassword("123456", "panicgo")
	fmt.Println(hashPwd)
}

func TestVerifyPassword(t *testing.T) {
	var user User
	user.Password = user.EncryptPassword("123456", user.Salt())
	fmt.Println(user.Password)
	correct := user.VerifyPassword("123456")
	assert.True(t, correct)
}

func TestPanicModels_CreateUser(t *testing.T) {
	var pm PanicModel
	pm.DB = database.GetTestDB()

	var user User
	user.Name = "yiming"
	user.Email = "yiming@qq.com"
	user.Password = user.EncryptPassword("123456", "hello world")

	user, err := pm.CreateUser(user)
	assert.Nil(t, nil, err)
	t.Log(user)
}

func TestPanicModels_QueryUserById(t *testing.T) {
	var pm PanicModel
	pm.DB = database.GetTestDB()

	var data User
	data.Name = "yiming"
	data.Email = "yiming@qq.com"
	data.Password = data.EncryptPassword("123456", "hello world")

	user, err := pm.CreateUser(data)
	assert.Nil(t, nil, err)
	t.Log(user)

	result, err := pm.QueryUserById(user.ID)
	assert.Nil(t, nil, err)
	assert.Equal(t, user.ID, result.ID)
}

func TestPanicModels_UpdateUserInfo(t *testing.T) {
	var pm PanicModel
	pm.DB = database.GetTestDB()

	var data User
	data.Name = "yiming"
	data.Email = "yiming@qq.com"
	data.Password = data.EncryptPassword("123456", "hello world")

	user, _ := pm.CreateUser(data)
	t.Log(user)

	info := UserInfo{"updatedemail@go.com", 1, "Chengdu", "Golang", "18012345678", "avatarurl.jpg", "coverurl.jpg"}

	err := pm.UpdateUserInfo(user.ID, info)
	assert.Nil(t, nil, err)

	result, err := pm.QueryUserById(user.ID)
	t.Log(result)
	assert.Nil(t, nil, err)
	assert.Equal(t, info.Email, result.Email)
	assert.Equal(t, info.Sex, result.Sex)
	assert.Equal(t, info.Location, result.Location)
	assert.Equal(t, info.Introduce, result.Introduce)
	assert.Equal(t, info.Phone, result.Phone)
	assert.Equal(t, info.AvatarURL, result.AvatarURL)
	assert.Equal(t, info.CoverURL, result.CoverURL)
}

func TestPanicModels_DeleteUserById(t *testing.T) {
	var pm PanicModel
	pm.DB = database.GetTestDB()

	var data User
	data.Name = "yiming"
	data.Email = "yiming@qq.com"
	data.Password = data.EncryptPassword("123456", "hello world")

	user, _ := pm.CreateUser(data)
	t.Logf("%+v", user)

	err := pm.DeleteUserById(user.ID)
	assert.Nil(t, nil, err)

	result, err := pm.QueryUserById(user.ID)
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), result.ID)
}
