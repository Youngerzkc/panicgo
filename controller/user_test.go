package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitschain/panicgo/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	regInfo := model.AccountInfo{"leonzhao", "yimingdream@gmail.com", "123456"}

	bufStr, _ := json.Marshal(regInfo)
	c.Request, _ = http.NewRequest("POST", "/user/signup", bytes.NewBufferString(string(bufStr)))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	var pc *PanicController
	pc = GetTestController()
	pc.Signup(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.HeaderMap.Get("Content-Type"))

	var user model.User
	json.Unmarshal(w.Body.Bytes(), &user)

	fmt.Println(user)
	assert.Equal(t, "leonzhao", user.Name)
	assert.Equal(t, uint(1), user.Role)
	assert.Equal(t, uint(1), user.Status)
}
