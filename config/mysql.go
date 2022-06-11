package config

import (
	"github.com/mittacy/gin-toy/core/eMysql"
	"github.com/spf13/viper"
)

var mysqlConfigs map[string]eMysql.MultipleConf

func InitMysql() {
	mysqlConfigs = map[string]eMysql.MultipleConf{
		//"localhost": {
		//	Sources: []eMysql.Conf{
		//		{
		//			Host:     viper.GetString("MYSQL_LOCALHOST_RW_HOST"),
		//			Port:     viper.GetInt("MYSQL_LOCALHOST_RW_PORT"),
		//			Database: "db_name",
		//			User:     viper.GetString("MYSQL_LOCALHOST_RW_USER"),
		//			Password: viper.GetString("MYSQL_LOCALHOST_RW_PASSWORD"),
		//		},
		//	},
		//},
	}

	eMysql.Init(mysqlConfigs, eMysql.WithName("mysql"),
		eMysql.WithSlowThreshold(viper.GetDuration("GORM_SLOW_LOG_THRESHOLD")),
		eMysql.WithIgnoreRecordNotFound(true))
}
