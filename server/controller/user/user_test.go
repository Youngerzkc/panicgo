package user

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/bitschain/panicgo/server/test"
	"github.com/stretchr/testify/assert"
	"github.com/bitschain/panicgo/server/model"
	"fmt"
)

func TestSignup(t *testing.T) {
	test.InitTestDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	regInfo := SignUpInfo{"leonzhao", "yimingdream@gmail.com", "123456"}

	bufStr, _:= json.Marshal(regInfo)
	c.Request, _ = http.NewRequest("POST", "/user/signup", bytes.NewBufferString(string(bufStr)))
	c.Request.Header.Add("Content-Type", gin.MIMEJSON)

	Signup(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.HeaderMap.Get("Content-Type"))

	var user model.User
	json.Unmarshal(w.Body.Bytes(), &user)

	fmt.Println(user)
	assert.Equal(t, "leonzhao", user.Name)
	assert.Equal(t, uint(1), user.Role)
	assert.Equal(t, uint(1), user.Status)
}