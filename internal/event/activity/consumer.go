package activity

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/zhuguangfeng/go-chat/internal/repository/dao"
	"github.com/zhuguangfeng/go-chat/model"
	"github.com/zhuguangfeng/go-chat/pkg/logger"
	"github.com/zhuguangfeng/go-chat/pkg/saramax"
	"time"
)

const SyncActivityConsumerGroupName = "sync_activity"

type ActivityConsumer struct {
	client        sarama.Client
	logger        logger.Logger
	activityEsDao dao.ActivityEsDao
}

func NewActivityConsumer(client sarama.Client, logger logger.Logger, activityEsDao dao.ActivityEsDao) *ActivityConsumer {
	return &ActivityConsumer{
		client:        client,
		logger:        logger,
		activityEsDao: activityEsDao,
	}
}

func (a *ActivityConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient(SyncActivityConsumerGroupName, a.client)
	if err != nil {
		return err
	}
	go func() {
		err := cg.Consume(context.Background(), []string{TopicSyncActivity}, saramax.NewHandler[ActivityEvent](a.logger, a.Consumer))
		if err != nil {
			a.logger.Error("推出了消费 循环异常", logger.String("ConsumerGroupName", SyncActivityConsumerGroupName), logger.Error(err))
		}
	}()
	return err
}

func (a *ActivityConsumer) Consumer(sg *sarama.ConsumerMessage, activity ActivityEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return a.activityEsDao.InputActivity(ctx, a.toEsEntity(activity))
}

func (a *ActivityConsumer) toEsEntity(activity ActivityEvent) model.ActivityEs {
	return model.ActivityEs{
		ID:                  activity.ID,
		SponsorID:           activity.SponsorID,
		Title:               activity.Title,
		Desc:                activity.Desc,
		Media:               activity.Media,
		AgeRestrict:         activity.AgeRestrict,
		GenderRestrict:      activity.GenderRestrict,
		CostRestrict:        activity.CostRestrict,
		Visibility:          activity.Visibility,
		MaxPeopleNumber:     activity.MaxPeopleNumber,
		CurrentPeopleNumber: activity.CurrentPeopleNumber,
		Address:             activity.Address,
		Category:            activity.Category,
		StartTime:           activity.StartTime,
		DeadlineTime:        activity.DeadlineTime,
		Status:              activity.Status,
		CreatedTime:         activity.CreatedTime,
		UpdatedTime:         activity.UpdatedTime,
	}
}
