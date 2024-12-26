package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViperWatch()
	app := InitWebServer()

	for _, consumer := range app.Consumers {
		err := consumer.Start()
		if err != nil {
			panic(err)
		}
	}

	err := app.Server.Run(":8888")
	if err != nil {
		panic(err)
	}
}

func initViperWatch() {
	cfile := pflag.String("config", "./config/dev.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	viper.WatchConfig()

	//读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
