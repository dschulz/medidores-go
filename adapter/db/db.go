package db

import (
	"database/sql"
	"medidores/adapter/common"
	"medidores/config"
)

func New(conf *config.Conf) (*sql.DB, error) {

	pgcnf := &common.PostgresConfig{
		Host:     conf.Db.Host,
		Port:     conf.Db.Port,
		User:     conf.Db.Username,
		DBName:   conf.Db.DbName,
		Password: conf.Db.Password,
		SSLMode:  conf.Db.SSLMode,
	}

	pgConnString := pgcnf.ConnString()

	return sql.Open("postgres", pgConnString)

}