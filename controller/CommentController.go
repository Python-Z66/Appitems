package controller

import (
	"Appitems/common"
	"Appitems/dao"
	"Appitems/models"
	"Appitems/respones"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

var DB = dao.GetDB()

// 评论功能

func NewComment(c *gin.Context) {
	var (
		Comment       models.BlogComment
		CreatedatTime time.Time
	)
	DB = dao.GetDB()
	DB = dao.DatabaseConnection(DB)
	defer DB.Close()
	username, whether := c.Get("username")
	if !whether {
		respones.Fail(c, common.SystemError, "系统错误", nil)
		return
	}
	if err := c.ShouldBindJSON(&Comment); err != nil {
		panic(err)
	}
	content := Comment.Content

	title := Comment.Title
	//fmt.Println(title)
	times := Comment.CreatedAt

	Comment = models.BlogComment{
		Name:      username.(string),
		Content:   content,
		CreatedAt: times,
		Title:     title,
	}

	DB.Create(&Comment)
	respones.Success(c, common.CommentOK, "评论成功", gin.H{
		"id":   Comment.ArticleID,
		"time": CreatedatTime,
	})
}

// 评论信息返回
func CommentInfo(c *gin.Context) {
	DB = dao.GetDB()
	DB = dao.DatabaseConnection(DB)
	var comment []models.BlogComment
	defer DB.Close()
	if err := DB.Find(&comment).Error; err != nil {
		respones.Response(c, http.StatusInternalServerError, common.CommentBad, nil, "评论提取失败")
	}
	respones.Success(c, common.CommentOK, "查询成功", gin.H{
		"all_comments": comment,
	})
}

/*

文章管理
发布
删除
更新
*/

// 发布文章

func NewArticle(c *gin.Context) {

}
