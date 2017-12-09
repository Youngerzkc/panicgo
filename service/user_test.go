package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicService_CreateNewUser(t *testing.T) {
	s := GetTestService()

	data := AccountInfo{"leonzhao", "leonzhao@qq.com", "123456"}

	user, err := s.CreateNewUser(data)
	assert.Nil(t, err)
	t.Logf("%+v", user)
}
