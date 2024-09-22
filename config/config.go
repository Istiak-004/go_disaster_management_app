package config

import (
	"fmt"
	"log"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{
		DatabaseURL: "host=localhost user=postgres password=postgres dbname=dmb port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		Port:        "5000",
		JWTSecret:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}

	fmt.Println("env value is : ", AppConfig)

	if AppConfig.DatabaseURL == "" {
		log.Fatal("Database URL is required!")
	}
}
