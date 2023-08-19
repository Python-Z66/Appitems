package common

// 错误码定义

const (
	UserNotExist          = 1001 // 用户不存在
	UserPasswordError     = 1002 // 用户密码错误
	UserDuplicateError    = 1003 // 用户重复注册
	UserEmailBindedError  = 1004 // 用户email已绑定
	UserRegister          = 2001 // 用户注册成功
	CommentOK             = 2003
	UserValueIsblank      = 1007 // 注册数据为空
	FoundUserOK           = 2002 // 查找用户信息成功
	UserLoginOK           = 2202 // 登录成功
	UserOutLoginOK        = 2204
	UserLoginBad          = 4001 // 用户登录失败
	CommentBad            = 4005 // 没有登录 不能评论
	CommentError          = 4006 //
	NewTokenError         = 4002 // token 生成错误
	UserUnauthorized      = 4004 // 权限不足
	HashUserPasswordError = 5001 // 用户密码加密失败
	NewUserBad            = 5000 // 创建用户失败
	SystemError           = 5005
)
