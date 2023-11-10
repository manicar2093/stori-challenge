package config

import (
	"github.com/spf13/viper"
)

var Instance = &config{}

type config struct {
	Environment string
	DatabaseURL string
	PageSize    int64
}

func init() {
	viper.AutomaticEnv()

	Instance.Environment = viper.GetString("ENVIRONMENT")
	Instance.DatabaseURL = viper.GetString("DATABASE_URL")
	Instance.PageSize = viper.GetInt64("PAGE_SIZE")
}
