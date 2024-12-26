package model

import (
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

const TableNameActivity = "activity"

type ActivityStatus uint

const (
	ActivityStatusUnknown        ActivityStatus = iota
	ActivityStatusPendingReview                 //待审核
	ActivityStatusReviewFailed                  //审核未通过
	ActivityStatusSignup                        //报名中==审核通过
	ActivityStatusDeadlineSignup                //截止报名
	ActivityStatusStart                         //进行中
	ActivityStatusEnd                           //活动结束
)

func (s ActivityStatus) Uint() uint {
	return uint(s)
}

type Activity struct {
	Base
	SponsorID           int64   `gorm:"column:sponsor_id;type:bigint;index:idx_sponsor_id;not null;comment:发起人ID" json:"sponsorId"`
	GroupID             int64   `gorm:"column:group_id;type:bigint;not null;comment:活动群聊id" json:"groupId"`
	Title               string  `gorm:"column:title;type:varchar(255);index:idx_title_status_gender_age_cost_visibility;not null;default:'';comment:活动标题" json:"title"`
	Status              uint    `gorm:"column:status;type:tinyint;index:idx_title_status_gender_age_cost_visibility;not null;default:0;comment:活动状态" json:"status"`
	Desc                string  `gorm:"column:desc;type:text;not null;comment:活动描述" json:"desc"`
	Media               Strings `gorm:"column:media;type:text;not null;comment:活动图片或视频" json:"media"`
	GenderRestrict      uint    `gorm:"column:gender_restrict;type:tinyint;index:idx_title_status_gender_age_cost_visibility;not null;default:0;comment:性别限制 男|女|不限" json:"genderRestrict"`
	AgeRestrict         uint    `gorm:"column:age_restrict;type:tinyint;index:idx_title_status_gender_age_cost_visibility;not null;default:0;comment:最大年龄限制" json:"ageRestrict"`
	CostRestrict        uint    `gorm:"column:cost_restrict;type:tinyint;index:idx_title_status_gender_age_cost_visibility;not null;default:0;comment:费用支付方式" json:"CostRestrict"`
	Visibility          uint    `gorm:"column:visibility;type:tinyint;index:idx_title_status_gender_age_cost_visibility;not null;default:0;comment:报名可见度" json:"visibility"`
	MaxPeopleNumber     int64   `gorm:"column:max_people_number;type:tinyint;not null;default:0;comment:最大报名人数" json:"maxPeopleNumber"`
	CurrentPeopleNumber int64   `gorm:"column:current_people_number;type:tinyint;not null;default:0;comment:当前报名人数" json:"CurrentPeopleNumber"`
	Address             string  `gorm:"column:address;type:varchar(255);not null;comment:获取地点" json:"address"`
	StartTime           uint    `gorm:"column:start_time;type:int;not null;default:0;comment:活动开始时间" json:"startTime"`
	DeadlineTime        uint    `gorm:"column:deadline_time;type:int;not null;default:0;comment:活动报名截止时间" json:"deadlineTime"`
	Category            uint    `gorm:"column:category;type:int;not null;default:0;comment:活动类型" json:"category"`
}

func (a *Activity) ToDto() dtov1.Activity {
	return dtov1.Activity{
		ID: a.ID,
		Sponsor: dtov1.User{
			ID: a.SponsorID,
		},
		Group: dtov1.Group{
			ID: a.GroupID,
		},
		Title:               a.Title,
		Status:              a.Status,
		Desc:                a.Desc,
		Media:               a.Media,
		GenderRestrict:      a.GenderRestrict,
		AgeRestrict:         a.AgeRestrict,
		CostRestrict:        a.CostRestrict,
		Visibility:          a.Visibility,
		MaxPeopleNumber:     a.MaxPeopleNumber,
		CurrentPeopleNumber: a.CurrentPeopleNumber,
		Address:             a.Address,
		StartTime:           utils.ToTimeString(a.StartTime),
		DeadlineTime:        utils.ToTimeString(a.DeadlineTime),
		Category:            a.Category,
	}
}

func (a *Activity) TableName() string {
	return TableNameActivity
}
