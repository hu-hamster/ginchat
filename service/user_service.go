package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取成功",
		"data":    data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	user.Salt = fmt.Sprintf("%6d", rand.Int31())

	user2 := models.FindUserByName(user.Name)
	if user2.Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名已注册！",
		})
		return
	}
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "两次密码不一致！",
		})
		return
	}
	user.Password = utils.MakePassword(password, user.Salt)
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功新增用户！",
	})
}

// FindUserByNameAndPwd
// @Summary 单用户信息
// @Tags 用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "登录成功",
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data := models.FindUserByNameAndPwd(name, pwd)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    data,
	})
}

// CreateUser
// @Summary 删除用户
// @Tags 用户
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功删除用户！",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改参数不匹配",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功修改用户！",
	})
}
