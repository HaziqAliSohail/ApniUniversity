package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"

	ServerHost = "server.host"
	ServerPort = "server.port"
)

func init() {
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")

	// defaults Database
	viper.SetDefault(DbName, "ApniUniversity")
	viper.SetDefault(DbHost, "127.0.0.1")
	viper.SetDefault(DbPort, "27017")

	//Server
	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8081")
}
