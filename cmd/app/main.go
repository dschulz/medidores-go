package main

import (
	"fmt"
	"github.com/gorilla/mux"
	dbConn "medidores/adapter/gorm"
	"medidores/app/app"
	"medidores/app/router"
	"medidores/config"
	"net/http"

	lr "medidores/util/logger"
	vr "medidores/util/validator"
)

func inicio (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<html><h1>Okay</h1></html>")

}


func main (){
	appConf := config.AppConfig()

	r := mux.NewRouter()

	db, _ := dbConn.New(appConf)

	// TODO: Arreglar este desastre

	r.HandleFunc("/", inicio)

	logger := lr.New(appConf.Debug)

	validator := vr.New()

	application := app.New(logger, db, validator)

	appRouter := router.New(application)

	address := ":3000"

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		//requestlog.Fatal().Err(err).Msg("Server startup failed")
		fmt.Println(err)
	}
}