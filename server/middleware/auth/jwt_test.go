package auth

import (
	"testing"
	"fmt"
)

func TestJWT_CreateToken(t *testing.T) {
	j := NewJWT()
	claims := CustomClaims{ID: 1, Name: "leonzhao", Email: "yimingdream@gmail.com"}
	token, _:= j.CreateToken(claims)
	fmt.Println(token)
}
