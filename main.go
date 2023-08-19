package main

import (
	"Appitems/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// gin+gorm+jwt 项目 ----> 基础企业项目
/*
jwt  go get -u github.com/dgrijalva/jwt-go  ---> 加密模块
gin  go get -u github.com/gin-gonic/gin ---> web框架
cors go get -u github.com/gin-contrib/cors  --> 解决跨域问题
gorm go get -u github.com/jinzhu/gorm ------>jorm
session go get -u github.com/gin-contrib/sessions
viper go get -u github.com/spf13/viper ----->读取并解析配置文件（例如 JSON、YAML、TOML、INI等）中的配置项；
	读取并解析环境变量、命令行参数等外部数据源中的配置项；
	提供简洁的调用方式来获取和设置配置项；
	支持在多个配置文件中层级继承和覆盖配置项。

*/
func main() {
	InitConfig()
	r := gin.Default()
	// 解决跨域问题  可以加header 来 但是这里我为了方便
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{
		"Authorization",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Methods",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Accept-Language",
		"Accept-Encoding",
		"Cookie",
	}
	config.AllowOrigins = []string{"http://127.0.0.1:5500", "http://localhost:63343", "http://localhost:63342"} // 允许的域名或 IP 地址
	config.AllowCredentials = true
	config.AddAllowMethods("OPTIONS")
	r.Use(cors.New(config))
	//r.Use(cors.Default())
	stroe := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", stroe))
	r = routers.Router(r)
	port := viper.GetString("server.prot")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// 读取配置文件

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err = viper.ReadInConfig()
	if err != nil {

	}
}
