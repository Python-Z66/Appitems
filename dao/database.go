package dao

import (
	"Appitems/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// 数据库处理
var db *gorm.DB

// 连接数据库
func DatabaseConnection(DB *gorm.DB) *gorm.DB {
	driverName := viper.GetString("datasource.driverName") // 数据库类型
	username := viper.GetString("datasource.username")     // 账号
	password := viper.GetString("datasource.password")     // 密码
	host := viper.GetString("datasource.host")
	prot := viper.GetString("datasource.prot")       // ip和端口
	dbname := viper.GetString("datasource.database") // 数据库名称
	charset := viper.GetString("datasource.charset")
	aigis := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%smb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		prot,
		dbname,
		charset,
	)
	var err error
	DB, err = gorm.Open(driverName, aigis)
	if err != nil {
		panic(err)
	}
	// 自定义表名
	//DB.Table("userdatabase").CreateTable(&models.UserDatabase{})
	DB.AutoMigrate(
		&models.UserDatabase{},
		&models.UserDto{},
		&models.IpRecord{},
		&models.BlogComment{},
		&models.BlogArticle{},
		&models.Claims{},
	)
	return DB
}

func GetDB() *gorm.DB {
	return db
}
