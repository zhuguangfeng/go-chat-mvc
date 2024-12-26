package common

import "strings"

const (
	StatusOk  int = 200
	StatusErr int = -1
)

type ErrCode string

func (c ErrCode) GetCodeMsg() (string, string) {
	str := string(c)
	index := strings.Index(str, ":")
	return str[:index], str[index+1:]
}

var (
	SystemError       ErrCode = "GoChat.Server:系统错误"
	BadRequestInvalid ErrCode = "GoChat:BadRequestInvalid:请求参数有误"

	UserPhoneNotFount   ErrCode = "GoChat.User.PhoneNotFound:用户手机号码未注册"
	UserPhoneExists     ErrCode = "GoChat.User.PhoneExists:手机号码已存在"
	UserPasswordInvalid ErrCode = "GoChat.User.PasswordInvalid:密码错误"
	UserDeregister      ErrCode = "GoChat.User.Deregister:用户已注销"

	ActivityNotFound ErrCode = "GoChat.Activity.NotFound:活动不存在"

	ReviewedNotAudit ErrCode = "GoChat.Review.ReviewedNotAudit:审核过的不能再次审核"
)
