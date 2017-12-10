package model

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

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

type AccountInfo struct {
	Name     string `json:"name" valid:"runelength(4|20)"`
	Email    string `json:"email" valid:"email,runelength(5|50)"`
	Password string `json:"password" valid:"runelength(6|20)"`
}

type UserInfo struct {
	Email     string `json:"email"`
	Sex       uint   `json:"sex"`
	Location  string `json:"location"`
	Introduce string `json:"introduce"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatarURL"`
	CoverURL  string `json:"coverURL"`
}

func (user *User) Salt() string {
	var userSalt string
	if user.Password == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = user.Password[0:10]
	}
	return userSalt
}

func (user *User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + "pwdsalt" //todo 这里hard code ，需要改为配置文件
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return hash
}

func (user *User) VerifyPassword(password string) bool {
	if password == "" || user.Password == "" {
		return false
	}

	ep := user.EncryptPassword(password, user.Salt())
	return ep == user.Password
}

func (pm *PanicModel) CreateUser(user User) (User, error) {
	if err := pm.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (pm *PanicModel) QueryUserById(id uint) (User, error) {
	var user User
	if err := pm.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (pm *PanicModel) UpdateUserInfo(id uint, info UserInfo) error {
	var user User
	if err := pm.DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := pm.DB.Model(&user).Where("id = ?", id).Updates(info).Error; err != nil {
		return err
	}
	return nil
}

func (pm *PanicModel) DeleteUserById(id uint) error {
	var user User
	if err := pm.DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := pm.DB.Model(&user).Where("id = ?", id).Update("DeletedAt", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
