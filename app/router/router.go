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

	// Routes for health
	r.HandleFunc("/health/liveness", app.HandleLive)
	r.Handle( "/health/readiness", requestlog.NewHandler(a.HandleReady, l)).Methods("GET")

    r.Use(middleware.ContentTypeJson)  // TODO: llevar abajo esto

	//// Routes for APIs
	//r.Route("/api/v1", func(r chi.Router) {
	//	r.Use(middleware.ContentTypeJson)
	//
	//	// Routes for books
	//	r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBooks, l))
	//	r.Method("POST", "/books", requestlog.NewHandler(a.HandleCreateBook, l))
	//	r.Method("GET", "/books/{id}", requestlog.NewHandler(a.HandleReadBook, l))
	//	r.Method("PUT", "/books/{id}", requestlog.NewHandler(a.HandleUpdateBook, l))
	//	r.Method("DELETE", "/books/{id}", requestlog.NewHandler(a.HandleDeleteBook, l))
	//})
	//
	//r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}