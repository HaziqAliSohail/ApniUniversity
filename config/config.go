package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
)

func init() {
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")

	// defaults
	viper.SetDefault(DbName, "ApniUniversity")
	viper.SetDefault(DbHost, "127.0.0.1")
	viper.SetDefault(DbPort, "27017")
}
