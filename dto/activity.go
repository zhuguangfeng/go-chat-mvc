package dto

import (
	"time"
)

type CreateActivityReq struct {
	SponsorID       int64    `json:"sponsorId"`       //发起人ID
	Title           string   `json:"title"`           //标题
	Desc            string   `json:"desc"`            //描述
	Media           []string `json:"media"`           //资源
	GenderRestrict  uint     `json:"genderRestrict"`  //性别限制
	AgeRestrict     uint     `json:"ageRestrict"`     //最大年龄
	CostRestrict    uint     `json:"costRestrict"`    //费用方式
	Visibility      uint     `json:"visibility"`      //活动可见度
	MaxPeopleNumber int64    `json:"maxPeopleNumber"` //最多人数
	Address         string   `json:"address"`         //地址
	StartTime       string   `json:"startTime"`       //开始时间
	DeadlineTime    string   `json:"deadlineTime"`    //结束报名时间
	Category        uint     `json:"category"`        //类型
}

func (c CreateActivityReq) StartTimeToTimestamp() (uint, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, c.StartTime)
	if err != nil {
		return 0, err
	}
	return uint(parsedTime.Unix()), nil
}

func (c CreateActivityReq) DeadlineTimeToTimestamp() (uint, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, c.DeadlineTime)
	if err != nil {
		return 0, err
	}
	return uint(parsedTime.Unix()), nil
}

type ActivitySearchReq struct {
	BaseListReq
	AgeRestrict    uint   `json:"ageRestrict"`
	GenderRestrict uint   `json:"genderRestrict"`
	CostRestrict   uint   `json:"CostRestrict"`
	Visibility     uint   `json:"visibility"`
	Address        string `json:"address"`
	Category       uint   `json:"category"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"EndTime"`
	Status         uint   `json:"status"`
}

type Activity struct {
	ID                  int64    `json:"id,omitempty"`
	Sponsor             User     `json:"sponsor,omitempty"`
	Group               Group    `json:"group,omitempty"`
	Title               string   `json:"title,omitempty"`
	Status              uint     `json:"status,omitempty"`
	Desc                string   `json:"desc,omitempty"`
	Media               []string `json:"media,omitempty"`
	GenderRestrict      uint     `json:"genderRestrict,omitempty"`
	AgeRestrict         uint     `json:"ageRestrict,omitempty"`
	CostRestrict        uint     `json:"costRestrict,omitempty"`
	Visibility          uint     `json:"visibility,omitempty"`
	MaxPeopleNumber     int64    `json:"maxPeopleNumber,omitempty"`
	CurrentPeopleNumber int64    `json:"currentPeopleNumber,omitempty"`
	Address             string   `json:"address,omitempty"`
	StartTime           string   `json:"startTime,omitempty"`
	DeadlineTime        string   `json:"deadlineTime,omitempty"`
	Category            uint     `json:"category,omitempty"`
}
