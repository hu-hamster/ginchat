package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config *viper.Viper
	DB     *gorm.DB
)

func InitConfig() {
	Config = viper.New()
	Config.AddConfigPath("config")
	Config.SetConfigName("app")
	Config.SetConfigType("yml")
	if err := Config.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func InitMySQL() {
	//自定义日志打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //彩色
		},
	)
	user := Config.GetString("mysql.user")
	password := Config.GetString("mysql.password")
	address := Config.GetString("mysql.address")
	port := Config.GetString("mysql.port")
	database := Config.GetString("mysql.database")
	dns := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
}
