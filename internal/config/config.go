package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `yaml:"logLevel"`
}

func Read() Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("[Error] Reading Viper Config")
		panic(err)
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		fmt.Println("[Error] Unmarshaling Viper Config")
		panic(err)
	}

	// Config 파일 변경 감지
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed. Operation: %v\n", e.Op.String())
	})

	return config
}
