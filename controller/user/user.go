package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leonzhao/panicgo/model"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type AccountInfo struct {
	Name     string `json:"name" valid:"runelength(4|20)"`
	Email    string `json:"email" valid:"email,runelength(5|50)"`
	Password string `json:"password" valid:"runelength(6|20)"`
}

type PersonalInfo struct {
	Email     string `json:"email"`
	Sex       uint   `json:"sex"`
	Location  string `json:"location"`
	Introduce string `json:"introduce"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatarURL"`
	CoverURL  string `json:"coverURL"`
}

// Signup 用户注册
func Signup(c *gin.Context) {
	var data AccountInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
	}

	var user model.User
	user.Name = data.Name
	user.Email = data.Email
	user.Password = user.EncryptPassword(data.Password, user.Salt())
	user.Role = model.UserRoleNormal
	user.Status = model.UserStatusInActive
	user.AvatarURL = "/images/avatar" + strconv.Itoa(rand.Intn(2)) + ".png"

	if err := model.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "invalid"})
	}
	c.JSON(http.StatusOK, user)
}

// UpdateInfo 更新用户信息
func UpdateInfo(c *gin.Context) {
	var data PersonalInfo
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid parameters"})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid parameters"})
	}

	jsonData, _ := json.Marshal(data)
	log.Printf("%v", string(jsonData))
	var user model.User
	model.DB.First(&user, id)

	//model.DB.Model(&user).Update("email", data.Email)
	//model.DB.Model(&user).Update(data)
	model.DB.Model(&user).Where("id = ?", id).Updates(data)

	c.JSON(http.StatusOK, user)
}

// PublicInfo 用户的公开信息
func PublicInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "invalid parameters"})
	}

	var user model.User
	log.Println(id)
	if err := model.DB.Model(&user).Where("id = ?", 1).First(&user).Error; err != nil {
		panic(err)
	}

	log.Println(user)
	c.JSON(http.StatusOK, user)
}

// Signin 用户登录
func Signin(c *gin.Context) {

}

// Signout 用户退出
func Signout(c *gin.Context) {

}

// ActiveAccount 激活账号
func ActiveAccount(c *gin.Context) {

}

//　发送重置密码邮件
func SendResetPasswordMail(c *gin.Context) {

}

// 验证重置密码链接
func VerifyResetPasswordLink(c *gin.Context) {

}

// UpdatePassword 更新用户密码
func UpdatePassword(c *gin.Context) {

}

// ResetPassword 重置用户密码
func ResetPassword(c *gin.Context) {

}

// SecretInfo 用户的隐私信息
func SecretInfo(c *gin.Context) {

}

// InfoDetail 用户信息详情
func InfoDetail(c *gin.Context) {

}

// AllList 所有用户列表
func AllList(c *gin.Context) {

}

// topN 前N名用户
func topN(n int, c *gin.Context) {

}

// TopTen 前10名用户
func TopTen(c *gin.Context) {

}

// TopHundred 前100名用户
func TopHundred(c *gin.Context) {

}

// UpdateAvatar 更新用户头像
func UpdateAvatar(c *gin.Context) {

}
