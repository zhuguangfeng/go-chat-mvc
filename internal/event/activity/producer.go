package activity

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
)

const TopicSyncActivity = "sync_activity_event"

type Producer interface {
	ProducerSyncActivityEvent(ctx context.Context, activity ActivityEvent) error
}

type ActivityProducer struct {
	logger   logger.Logger
	producer sarama.SyncProducer
}

func NewProducer(logger logger.Logger, producer sarama.SyncProducer) Producer {
	return &ActivityProducer{
		logger:   logger,
		producer: producer,
	}
}

func (a *ActivityProducer) ProducerSyncActivityEvent(ctx context.Context, activity ActivityEvent) error {
	val, err := json.Marshal(activity)
	if err != nil {
		return err
	}
	partition, offset, err := a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: TopicSyncActivity,
		Value: sarama.StringEncoder(val),
	})
	if err != nil {
		return err
	}
	a.logger.Info("发送同步活动时间成功", logger.Int32("partition", partition), logger.Int64("offset", offset), logger.Any("activity", activity))
	return nil
}

type ActivityEvent struct {
	ID                  int64    `json:"id"`
	SponsorID           int64    `json:"sponsorId"`
	Title               string   `json:"title"`
	Desc                string   `json:"desc"`
	Media               []string `json:"media"`
	AgeRestrict         uint     `json:"ageRestrict"`
	GenderRestrict      uint     `json:"genderRestrict"`
	CostRestrict        uint     `json:"costRestrict"`
	Visibility          uint     `json:"visibility"`
	MaxPeopleNumber     int64    `json:"maxPeopleNumber"`
	CurrentPeopleNumber int64    `json:"currentPeopleNumber"`
	Address             string   `json:"address"`
	Category            uint     `json:"category"`
	StartTime           uint     `json:"startTime"`
	DeadlineTime        uint     `json:"deadlineTime"`
	Status              uint     `json:"status"`
	CreatedTime         uint     `json:"createdTime"`
	UpdatedTime         uint     `json:"updatedTime"`
}
