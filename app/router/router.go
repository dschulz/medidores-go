package router

import (

	"github.com/gorilla/mux"

	"medidores/app/app"
	"medidores/app/requestlog"
	"medidores/app/router/middleware"
)

func New(a *app.App) *mux.Router {
	l := a.Logger()

    r := mux.NewRouter()

	// Health
	r.HandleFunc("/health/liveness", app.HandleLive)
	r.Handle( "/health/readiness", requestlog.NewHandler(a.HandleReady, l)).Methods("GET")

    sr := r.PathPrefix("/api/v1").Subrouter()
    sr.Use(middleware.ContentTypeJson)

    sr.Handle("/medidor", requestlog.NewHandler(a.HandleListarMedidores ,l)).Methods("GET")
	sr.Handle("/medidor", requestlog.NewHandler(a.HandleCrearMedidor ,l)).Methods("POST")
	sr.Handle("/medidor/{1}", requestlog.NewHandler(a.HandleBuscarMedidor ,l)).Methods("GET")
	sr.Handle("/medidor/{1}", requestlog.NewHandler(a.HandleActualizarMedidor ,l)).Methods("PUT")
	sr.Handle("/medidor/{1}", requestlog.NewHandler(a.HandleEliminarMedidor ,l)).Methods("DELETE")



	return r
}