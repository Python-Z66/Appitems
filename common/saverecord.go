package common

import (
	"Appitems/models"
	"Appitems/respones"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

// 管理员登陆IP处理 保存记录ip

/*
每个 IP 的错误次数和最后一次错误时间，并使用 GORM 存储到数据库中
*/
//var db = dao.GetDB()

func LimitAccessMiddleware(ip string, db *gorm.DB, c *gin.Context) {
	var logip models.IpRecord
	// 查询访问记录
	db.FirstOrCreate(&logip, models.IpRecord{IP: ip})
	now := time.Now()
	//.Format("2006-01-02 15:04:05")
	duration := now.Sub(logip.LastTime)
	if duration > Day {
		logip.ErrorCnt = 0
	}
	if logip.ErrorCnt > 3 {
		// 超过访问限制，拦截请求
		logip.StartTime = now
		logip.Duration = Day
		db.Save(&logip)
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "错误次数过多"})

		return
	}
	// 更新访问记录
	logip.IP = ip
	logip.ErrorCnt++
	logip.LastTime = now
	db.Save(&logip)
	respones.Response(c, http.StatusBadRequest, 233, nil, fmt.Sprintf("账号密码错误,还剩%d次禁止访问", 4-logip.ErrorCnt))
}
