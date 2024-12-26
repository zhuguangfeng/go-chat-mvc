package ioc

import (
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"github.com/zhuguangfeng/go-chat/internal/event/activity"
	"github.com/zhuguangfeng/go-chat/pkg/saramax"
)

func InitKafka() sarama.Client {
	type Config struct {
		Addr []string `yaml:"addr"`
	}
	var cfg Config
	err := viper.UnmarshalKey("kafka", &cfg)
	if err != nil {
		panic(err)
	}

	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true
	client, err := sarama.NewClient(cfg.Addr, scfg)
	if err != nil {
		panic(err)
	}
	return client
}

func NewConsumers(activityConsumer *activity.ActivityConsumer) []saramax.Consumer {
	return []saramax.Consumer{activityConsumer}
}

func InitSaramaSyncProducer(client sarama.Client) sarama.SyncProducer {
	p, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
	return p
}
