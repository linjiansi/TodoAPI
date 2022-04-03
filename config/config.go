package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	appConfig *config
	once sync.Once
)

type config struct {
	DBHost     string
	DBPort     string
	DBUserName string
	DBPassword string
	DBName     string
	Port       string
	SecretKey  string
}

func (config *config) DBSrc() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBUserName,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}

func Config() *config {
	once.Do(func() {
		appConfig = &config{
			DBHost:      envOrDefault("APP_DBHOST", "localhost"),
			DBPort:      envOrDefault("APPPORT", "3306"),
			DBUserName:  envOrDefault("APP_DBUSERNAME", "root"),
			DBPassword:  envOrDefault("APP_DBPASSWORD", "password"),
			DBName:      envOrDefault("APP_DBNAME", "todo_api_dev"),
			Port:        envOrDefault("APP_PORT", "8080"),
			SecretKey:   envOrDefault("APP_SECRETKEY", "JDJhJDEwJDE5clFvMFJJdkI1T0xBSlF6ci50Ny42am84Vjd4YXYwYVN0UHFuZTF4N1ZlbzFHdjBzd0dh"),
		}
	})

	return appConfig
}

func envOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}