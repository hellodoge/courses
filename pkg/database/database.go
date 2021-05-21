package database

import (
	"fmt"
)

type Config struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
}

func (cfg Config) GetMySQLDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
}
