package model

import (
	_ "embed"
	"github.com/zhuguangfeng/go-chat/model"
)

const ActivityIndexName = "activity"

var (
	//go:embed activity_index.json
	ActivityIndex string
)

type ActivityEs struct {
	ID                  int64    `json:"id"`
	SponsorID           int64    `json:"sponsorId"`
	GroupID             int64    `json:"groupId"`
	Title               string   `json:"title"`
	Desc                string   `json:"desc"`
	Media               []string `json:"media"`
	AgeRestrict         uint     `json:"ageRestrict"`
	GenderRestrict      uint     `json:"genderRestrict"`
	CostRestrict        uint     `json:"CostRestrict"`
	Visibility          uint     `json:"visibility"`
	MaxPeopleNumber     int64    `json:"maxPeopleNumber"`
	CurrentPeopleNumber int64    `json:"CurrentPeopleNumber"`
	Address             string   `json:"address"`
	Category            uint     `json:"category"`
	StartTime           uint     `json:"startTime"`
	DeadlineTime        uint     `json:"deadlineTime"`
	Status              uint     `json:"status"`
	CreatedTime         uint     `json:"createdTime"`
	UpdatedTime         uint     `json:"updatedTime"`
}

func (a *ActivityEs) ToModel() Activity {
	return model.Activity{
		Base: model.Base{
			ID:        a.ID,
			CreatedAt: a.CreatedTime,
			UpdatedAt: a.UpdatedTime,
		},
		Title:               a.Title,
		Desc:                a.Desc,
		Media:               a.Media,
		AgeRestrict:         a.AgeRestrict,
		GenderRestrict:      a.GenderRestrict,
		CostRestrict:        a.CostRestrict,
		Visibility:          a.Visibility,
		MaxPeopleNumber:     a.MaxPeopleNumber,
		CurrentPeopleNumber: a.CurrentPeopleNumber,
		Address:             a.Address,
		Category:            a.Category,
		StartTime:           a.StartTime,
		DeadlineTime:        a.DeadlineTime,
		Status:              a.Status,
	}
}
