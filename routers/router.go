package routers

import (
	"Appitems/controller"
	"Appitems/middleware"
	"github.com/gin-gonic/gin"
)

// 创建路由
// 路由函数
func Router(r *gin.Engine) *gin.Engine {
	// 用户注册
	r.POST("/api/auth/register", controller.Register)
	// 用户登录
	r.POST("/api/auth/login", controller.Login)
	// 用户信息
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	//管理员登陆
	r.POST("/admin/auth/login", controller.AdminLogin)
	// 删除用户登录信息
	r.DELETE("/api/auth/delete_user", controller.DeleteUserLogin)
	// 评论
	r.POST("/api/comment/new", middleware.CommentMiddleware(), controller.NewComment)
	// 评论信息
	r.GET("/api/comment/info", controller.CommentInfo)
	// 文章发布 管理员才能发布
	// -------------
	return r
}
