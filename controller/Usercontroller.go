package controller

import (
	"Appitems/common"
	"Appitems/dao"
	"Appitems/dto"
	"Appitems/models"
	"Appitems/respones"
	"Appitems/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

// 用户登陆注册逻辑
func Register(c *gin.Context) {
	var user models.UserDatabase
	DB = dao.GetDB()
	DB = dao.DatabaseConnection(DB)
	defer DB.Close()
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
	}
	email := user.Email
	password := user.Password
	username := user.Username
	DB.AutoMigrate(&models.UserDatabase{})
	var (
		countname  uint64
		countemail uint64
		err        error
	)
	if email != "" && password != "" && username != "" {
		err, err = DB.Table("user_databases").Where("username=?", username).Count(&countname).Error,
			DB.Table("user_databases").Where("email=?", email).Count(&countemail).Error
		if err != nil {
			respones.Response(c, http.StatusInternalServerError, common.NewUserBad, nil, "用户创建失败")
			return
		}
		if countname >= 1 {
			respones.Response(c, http.StatusBadRequest, common.UserDuplicateError, nil, "用户名被注册")
			return
		} else if countemail >= 1 {
			respones.Response(c, http.StatusBadRequest, common.UserEmailBindedError, nil, "邮箱被注册")
			return
		}
		//hashpassword, err := util.PasswordGend(password)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"code": common.HashUserPasswordError,
		//		"msg":  "加密失败",
		//	})
		//	return
		//}
		now := time.Now()
		newUser := models.UserDatabase{
			Username:  username,
			Email:     email,
			Password:  password,
			CreatedAt: now,
			UpdatedAt: now,
		}
		DB.Create(&newUser)
		respones.Success(c, common.UserRegister, "用户注册成功", nil)
		return
	}

	respones.Fail(c, common.UserValueIsblank, "不能为空", nil)
	return
}

// 用户登录

func Login(c *gin.Context) {
	var (
		user models.UserDatabase
	)
	DB = dao.GetDB()
	DB = dao.DatabaseConnection(DB)
	defer DB.Close()
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
	}
	username := user.Username
	password := user.Password
	hashpassword, err := util.PasswordGend(password)
	if err != nil {
		respones.Response(c, http.StatusInsufficientStorage, common.HashUserPasswordError, nil, "加密失败!!")
		return
	}
	DB.AutoMigrate(&models.UserDatabase{})
	IsFoundORNotFound := func(username, password string) bool {
		if restult := DB.Where("username=? AND password=?", username, password).First(&user); restult.Error == gorm.ErrRecordNotFound {
			// 没找到就是false
			return false
		} else if restult.Error != nil {
			panic(restult.Error)
		}
		return true
	}(username, password)
	if IsFoundORNotFound {
		se := sessions.Default(c)
		se.Set("hashpassword", hashpassword)
		se.Save()
		// 发放token
		token, errs := common.ReleaseTokenUser(user)
		if errs != nil {
			respones.Response(c, http.StatusInternalServerError, common.NewTokenError, nil, "系统异常!!")
			log.Println(errs)
			return
		}
		respones.Success(c, common.UserLoginOK, "登录成功", gin.H{"token": token, "user": user.Username})
		return
	}
	respones.Response(c, http.StatusBadRequest, common.UserLoginBad, nil, "账号或密码错误")
	return
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": common.FoundUserOK,
		"user": dto.ToUserDto(user.(models.UserDatabase)),
	})
}

// 清除用户登录信息

func DeleteUserLogin(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	keys := c.Keys
	for key := range keys {
		delete(keys, key)
	}
	// 2204
	respones.Success(c, common.UserOutLoginOK, "成功", nil)
}

// 管理员登录 --------->
func AdminLogin(c *gin.Context) {
	DB = dao.DatabaseConnection(DB)
	var AdminUser models.UserDatabase
	if err := c.ShouldBindJSON(&AdminUser); err != nil {
		panic(err)
	}
	adminusername := AdminUser.Username
	adminpassword := AdminUser.Password
	// 查询数据库
	adminfunc := func(adminusername, adminpassword string) bool {
		if resutlt := DB.Where("username=? AND password=?", adminusername, adminpassword).First(&AdminUser); resutlt.Error == gorm.ErrRecordNotFound {
			return false
		} else if resutlt.Error != nil {
			panic(resutlt.Error)
		}
		return true
	}(adminusername, adminpassword)

	if adminfunc {

		respones.Success(c, common.UserLoginOK, "登陆成功", nil)
		return
	}
	// 账号密码错误一个ip三次机会  一天解封
	ip := c.ClientIP()
	defer DB.Close()
	common.LimitAccessMiddleware(ip, DB, c)
}
