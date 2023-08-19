package middleware

import (
	"Appitems/common"
	"Appitems/dao"
	"Appitems/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strings"
)

// / 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// token 为空或者不是Bearer开头
		if token == "" || !strings.HasPrefix(token, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.UserUnauthorized,
				"massage": "权限不足",
			})
			c.Abort()
			return
		}
		token = token[7:]
		tok, claims, err := common.ParseToken(token)
		if err != nil || !tok.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.UserUnauthorized,
				"massage": "权限不足",
			})
			c.Abort()
			return
		}
		// token 通过了验证 获取token中的userid
		userid := claims.UserId
		DB := dao.GetDB()
		// 链接数据库
		DB = dao.DatabaseConnection(DB)
		defer DB.Close()
		var user models.UserDatabase
		DB.First(&user, userid)
		// 先验证用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.UserNotExist,
				"massage": "权限不足",
			})
			c.Abort()
			return
		}
		// 用户存在 把user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}

// 评论中间件
func CommentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// token 为空或者不是Bearer开头
		if token == "" || !strings.HasPrefix(token, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.CommentBad,
				"massage": "请先登录1",
			})
			c.Abort()
			return
		}
		token = token[7:]
		tok, claims, err := common.ParseToken(token)
		if err != nil || !tok.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.CommentBad,
				"massage": "请先登录2",
			})
			c.Abort()
			return
		}
		// token 通过了验证 获取token中的userid
		userid := claims.UserId
		DB := dao.GetDB()
		// 链接数据库
		DB = dao.DatabaseConnection(DB)
		defer DB.Close()
		var user models.UserDatabase
		DB.First(&user, userid)
		// 先验证用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    common.UserNotExist,
				"massage": "权限不足",
			})
			c.Abort()
			return
		}
		if user.Username == "blogadmin" {
			c.Set("username", "管理员")
			c.Next()
			return
		}
		c.Set("username", user.Username)
		c.Next()
	}
}

// 文章操作中间件
