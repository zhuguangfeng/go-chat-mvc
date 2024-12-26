package model

import (
	"encoding/json"
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

const TableNameUser = "user"

type UserStatus uint

func (u UserStatus) Uint() uint {
	return uint(u)
}

const (
	UserStatusUnknown            UserStatus = iota
	UserStatusNormal                        //正常
	UserStatusBan                           //账号封禁
	UserStatusProhibitedActivity            //禁止参加平台任何活动
)

type User struct {
	Base
	Username        string `gorm:"column:username;type:varchar(64);index:idx_username_gender;not null;default:'';comment:用户名称" json:"username"`
	Avatar          string `gorm:"column:avatar;type:text;comment:用户头像" json:"avatar"`
	BackgroundImage string `gorm:"column:backgroundImage;type:text;comment:背景图片" json:"backgroundImage"`
	Phone           string `gorm:"column:phone;type:char(11);uniqueIndex:udx_phone;not null;default:'';comment:手机号码" json:"phone"`
	Password        string `gorm:"column:password;type:varchar(128);not null;default:'';comment:用户密码" json:"password"`
	Age             uint   `gorm:"column:age;type:tinyint;not null;default:0;comment:年龄" json:"age"`
	Gender          uint   `gorm:"column:six;type:tinyint;index:idx_username_gender;not null;default:0;comment:性别" json:"gender"`
	IsRealName      bool   `gorm:"column:is_real_name;type:tinyint;not null;default:0;comment:是否实名认证" json:"isRealName"`
	IDCard          string `gorm:"column:id_card;type:char(18);not null;default:'';comment:身份证" json:"idCard"`
	Name            string `gorm:"column:name;type:varchar(32);not null;default:'';comment:真实姓名" json:"name"`
	LastLoginIp     string `gorm:"column:login_ip;type:varchar(32);not null;default:'';comment:最后登录ip" json:"loginIp"`
	LastLoginTime   uint   `gorm:"column:last_login_time;type:int;not null;default:0;comment:最后登录时间" json:"lastLoginTime"`
	Status          uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:账号状态" json:"status"`
}

func (u *User) TableName() string {
	return TableNameUser
}

func (u *User) ToDto() dtov1.User {
	return dtov1.User{
		ID:              u.ID,
		Username:        u.Username,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Gender:          u.Gender,
		Age:             u.Age,
		Status:          u.Status,
		CreatedTime:     utils.ToTimeString(u.CreatedAt),
	}
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
