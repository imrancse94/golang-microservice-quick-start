package config

import "os"

type SQLConf struct {
	Driver   string
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func GetSQL() *SQLConf {
	return &SQLConf{
		Driver:   "mysql",
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}
