package configs

import (
	"strconv"

	"github.com/spf13/viper"
)

var (
	TEMPORAL_CLIENT_URL string
	APP_API_VERSION     string
	APP_NAME            string
	APP_ENV             string
	APP_DEBUG_MODE      bool

	VERSION string

	DB_URL           string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_HOST          string
	DB_PORT          string
	DB_SSL_MODE      string
	DB_RUN_MIGRATION bool
	DB_RUN_SEEDER    bool
	DB_SCHEMA        string

	SERVER_HTTP_PORT string
)

func init() {
	NewViperConfig()

	var err error
	TEMPORAL_CLIENT_URL = viper.GetString("TEMPORAL_CLIENT_URL")

	APP_DEBUG_MODE, err = strconv.ParseBool(viper.GetString("APP_DEBUG_MODE"))
	if err != nil {
		APP_DEBUG_MODE = false
	}
	VERSION = viper.GetString("1.0.1")

	APP_API_VERSION = "v2"
	SERVER_HTTP_PORT = viper.GetString("APP_PORT")

	APP_NAME = viper.GetString("APP_NAME")
	APP_ENV = viper.GetString("APP_ENV")

	DB_URL = viper.GetString("DB_URL")

	DB_RUN_MIGRATION, err = strconv.ParseBool(viper.GetString("DB_RUN_MIGRATION"))
	if err != nil {
		DB_RUN_MIGRATION = false
	}

	DB_RUN_SEEDER, err = strconv.ParseBool(viper.GetString("DB_RUN_SEEDER"))
	if err != nil {
		DB_RUN_SEEDER = false
	}

}
