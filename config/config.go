package config

import (
	"taskema/datasource/mysql"
	"taskema/delivery/httpserver"
	"taskema/service/authservice"
)

type Config struct {
	Server httpserver.Config  `koanf:"server"`
	MySql  mysql.Config       `koanf:"mysql"`
	Auth   authservice.Config `koanf:"auth"`
}
