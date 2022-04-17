package config

import "go-grpc-server/util"

type Setting struct {
	Env        string
	DBName     string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
}

func NewSetting() Setting {
	return Setting{
		Env:        util.MustGetEnv("ENV", "development"),
		DBName:     util.MustGetEnv("DB_NAME", "go_grpc_server"),
		DBHost:     util.MustGetEnv("DB_HOST", "localhost"),
		DBPort:     util.MustAtoi(util.MustGetEnv("DB_PORT", "3306")),
		DBUser:     util.MustGetEnv("DB_USER", "root"),
		DBPassword: util.MustGetEnv("DB_PASSWORD", "1234"),
	}
}
