
package gorm

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"medidores/adapter/common"

	"medidores/config"
)




func New(conf *config.Conf) (*gorm.DB, error) {


	pgcnf := &common.PostgresConfig{
		Host:     conf.Db.Host,
		Port:     conf.Db.Port,
		User:     conf.Db.Username,
		DBName:   conf.Db.DbName,
		Password: conf.Db.Password,
		SSLMode: conf.Db.SSLMode,
	}

	pgConnString := pgcnf.ConnString()

	return gorm.Open("postgres", pgConnString)
}