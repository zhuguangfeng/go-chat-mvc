package ioc

import (
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"github.com/zhuguangfeng/go-chat/model"
	"time"
)

func InitEsClient() *elastic.Client {
	type Config struct {
		Url   string `yaml:"url"`
		Sniff bool   `json:"sniff"`
	}
	var cfg Config
	err := viper.UnmarshalKey("es", &cfg)
	if err != nil {
		panic(err)
	}

	const timeOut = time.Second * 10
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.Url),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheckInterval(timeOut),
	}

	client, err := elastic.NewClient(opts...)
	if err != nil {
		panic(err)
	}

	err = model.InitEs(client)
	if err != nil {
		panic(err)
	}

	return client
}
