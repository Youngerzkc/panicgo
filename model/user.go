package model

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

// User 用户
type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	Name         string     `json:"name"`
	Password     string     `json:"-"`
	Email        string     `json:"-"`
	Sex          uint       `json:"sex"`
	Location     string     `json:"location"`
	Introduce    string     `json:"introduce"`
	Phone        string     `json:"-"`
	Score        uint       `json:"score"`
	ArticleCount uint       `json:"articleCount"`
	CommentCount uint       `json:"commentCount"`
	CollectCount uint       `json:"collectCount"`
	Signature    string     `json:"signature"`
	Role         uint       `json:"role"`
	AvatarURL    string     `json:"avatarURL"`
	CoverURL     string     `json:"coverURL"`
	Status       uint       `json:"status"`
}

func (user User) Salt() string {
	var userSalt string
	if user.Password == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Password[0:10]
	}
	return userSalt
}

func (user User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + "pwdsalt" //todo 这里hard code ，需要改为配置文件
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return hash
}

func (user User) VerifyPassword(password string) bool {
	if password == "" || user.Password == "" {
		return false
	}

	ep := user.EncryptPassword(password, user.Salt())
	return ep == user.Password
 }

const (
	UserRoleNormal     = 1
	UserRoleEditor     = 2
	UserRoleAdmin      = 3
	UserRoleSuperAdmin = 4
	UserRoleCrawler    = 5
)

const (
	UserStatusInActive = 1
	UserStatusActived  = 2
	UserStatusFrozen   = 3
)

const (
	UserSexFemale = 0
	UserSexMale   = 1
)

const (
	MaxUserNameLen  = 20
	MinUserNameLen  = 4
	MaxPwdLen       = 20
	MinPwdLen       = 6
	MaxSignatureLen = 200
	MaxLocationLen  = 200
	MaxIntroduceLen = 500
)
