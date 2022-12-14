package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string `valid:"minstringlength(8)"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIP      string
	ClientPort    string
	Salt          string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64
	IsLogout      bool
	DeviceInfo    string
}

func (userBasic *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 0)
	utils.DB.Find(&data)
	return data
}

func FindUserByNameAndPwd(name, password string) *UserBasic {
	user := new(UserBasic)
	utils.DB.Where("name=? and password = ?", name, password).First(user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	fmt.Println(temp)
	utils.DB.Model(user).Where("id=?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) *UserBasic {
	user := new(UserBasic)
	utils.DB.Where("name=?", name).First(user)
	return user
}

func FindUserByPhone(phone string) *UserBasic {
	user := new(UserBasic)
	utils.DB.Where("phone=?", phone).First(user)
	return user
}

func FindUserByEmail(email string) *UserBasic {
	user := new(UserBasic)
	utils.DB.Where("email=?", email).First(user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(&UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email})
}
