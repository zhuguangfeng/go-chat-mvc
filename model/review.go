package model

import (
	dtov1 "github.com/zhuguangfeng/go-chat/dto"
	"github.com/zhuguangfeng/go-chat/pkg/utils"
)

const (
	TableNameReview = "review"
)
const (
	ReviewBizActivity = "activity"
)

type ReviewStatus uint

func (r ReviewStatus) Uint() uint {
	return uint(r)
}

const (
	ReviewStatusUnknown       ReviewStatus = iota
	ReviewStatusPendingReview              //待审核
	ReviewStatusReviewCancel               //审核取消
	ReviewStatusSuccess                    //审核通过
	ReviewStatusPass                       //审核拒绝
)

// Review 审核表
type Review struct {
	Base
	Biz        string `gorm:"column:biz;type:varchar(32);uniqueIndex:udx_biz_bizId;not null;default:'';comment:业务" json:"biz"`
	BizID      int64  `gorm:"column:biz_id;type:bigint;uniqueIndex:udx_biz_bizId;not null;default:0;comment:业务ID" json:"bizId"`
	SponsorID  int64  `gorm:"column:sponsor_id;type:bigint;not null;default:0;comment:发起人ID" json:"sponsorId"`
	ReviewerID int64  `gorm:"column:reviewer_id;type:bigint;not null;default:0;comment:审核人ID" json:"reviewerId"`
	Status     uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:审核状态" json:"status"`
	Opinion    string `gorm:"column:opinion;type:varchar(255);not null;default:'';comment:审核意见" json:"opinion"`
	ReviewTime uint   `gorm:"column:review_time;type:uint;not null;default:0;comment:审核时间"  json:"reviewTime"`
}

func (r Review) TableName() string {
	return TableNameReview
}

func (r Review) ToDto() dtov1.Review {
	return dtov1.Review{
		ID:     r.ID,
		Biz:    r.Biz,
		BizID:  r.BizID,
		Status: r.Status,
		Activity: dtov1.Activity{
			ID: r.BizID,
		},
		Reviewer: dtov1.User{
			ID: r.ReviewerID,
		},
		CreatedTime: utils.ToTimeString(r.CreatedAt),
		ReviewTime:  utils.ToTimeString(r.ReviewTime),
	}
}
