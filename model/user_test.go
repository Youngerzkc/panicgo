package model

import (
	"fmt"
	"testing"

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
