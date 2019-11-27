package common

import (
	"fmt"
	"strings"
)

type PostgresConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
	SSLMode string
}

func (c PostgresConfig) ConnString() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.User, c.DBName, c.Password, c.SSLMode))
	return builder.String()
}