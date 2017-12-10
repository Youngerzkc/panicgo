package service

import (
	"github.com/bitschain/panicgo/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicService_CreateNewUser(t *testing.T) {
	s := GetTestService()

	data := model.AccountInfo{"leonzhao", "leonzhao@qq.com", "123456"}

	user, err := s.CreateNewUser(data)
	assert.Nil(t, err)
	t.Logf("%+v", user)
}
