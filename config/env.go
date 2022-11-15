package config

import "github.com/ory/viper"

func BindEnvFor() {
	// 绑定所有环境变量
	viper.AutomaticEnv()
	// 绑定环境变量
	// viper.BindEnv("port", "PORT")
	// viper.BindEnv("workdir", "WORKDIR")
	// viper.BindEnv("db_type", "DB_TYPE")
	// viper.BindEnv("db_level", "DB_LEVEL")
	// viper.BindEnv("token", "TOKEN")
}
