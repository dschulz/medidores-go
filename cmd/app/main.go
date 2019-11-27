package main

import (
	"fmt"
	dbConn "medidores/adapter/gorm"
	"medidores/app/app"
	"medidores/app/router"
	"medidores/config"
	"net/http"

	lr "medidores/util/logger"
	vr "medidores/util/validator"
)


func main (){


	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	if appConf.Debug {
		db.LogMode(true)
	}

	validator := vr.New()

	application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Iniciando servicio en %v", address)


	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("No se pudo iniciar el servicio")
		fmt.Println(err)
	}
}