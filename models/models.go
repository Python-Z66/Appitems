package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 放模板

// 用户信息数据库
type UserDatabase struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// 验证ID 默认值为0 如果为会员或 管理员账号 改为888,222
	VerificationID uint   `gorm:"default:'0'" json:"verification-id"` // 没有为空的情况下
	Username       string `json:"username" gorm:"notnull;type:char(100);size:20"`
	Email          string `json:"email"`
	Password       string `json:"password" gorm:"notnull"`
}

//// 接收返回的管理员账号密码 并且每次登录都写入数据库 ----->我优化为userdatabase 就行了，
//type AdminDto struct {
//	ID       uint   `gorm:"primary_key"`
//	Name     string `json:"admin_name"`
//	Password string `json:"admin_password"`
//}

// 记录每个ip的错误时间和错误次数
type IpRecord struct {
	IP        string        `gorm:"primary_key"`
	ErrorCnt  uint          `gorm:"type:tinyint;not null;default:0"` // 错误次数
	StartTime time.Time     `gorm:"not null"`                        // 封禁开始时间
	Duration  time.Duration // 封禁时间
	LastTime  time.Time     `gorm:"not null"` //
}

// Token Claims

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 用户信息返回
type UserDto struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	VerificationID uint   `json:"verification_id"`
}

// 评论功能

type BlogComment struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"notnull"`
	ArticleID uint      `gorm:"AUTO_INCREMENT"` // 自增类型
	Content   string    `gorm:"notnull;size:200" json:"content"`
	CreatedAt time.Time `gorm:"notnull" json:"created_at"`
	Title     string    `gorm:"notnull" json:"title"` // 文章标题
	UpdatedAt time.Time
}

// 文章
type BlogArticle struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"notnull" json:"title"`      // 文章标题
	Content   string    `gorm:"notnull" json:"content"`    // 发送的文章
	CreatedAt time.Time `gorm:"notnull" json:"created_at"` // 发布时间
	UpdatedAt time.Time `gorm:"notnull" json:"updated_at"` // 更新时间
}
