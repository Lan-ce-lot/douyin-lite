package repository

import (
	"fmt"

	"github.com/houqingying/douyin-lite/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Setup(conf *config.DatabaseConfiguration) error {
	var err error
	username := conf.Username //账号
	password := conf.Password //密码
	host := conf.Host         //数据库地址，可以是Ip或者域名
	port := conf.Port         //数据库端口
	Dbname := conf.Dbname     //数据库名
	timeout := conf.Timeout   //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{}, &Comment{})
	return err
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return db
}
