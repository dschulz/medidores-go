package app

import "net/http"

func (app *App) HandleListarMedidores(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "nombre": "prueba" }`))
}

func (app *App) HandleBuscarMedidor(w http.ResponseWriter, r *http.Request) {
}

func (app *App) HandleCrearMedidor(w http.ResponseWriter, r *http.Request) {
}

func (app *App) HandleActualizarMedidor(w http.ResponseWriter, r *http.Request) {
}

func (app *App) HandleEliminarMedidor(w http.ResponseWriter, r *http.Request) {
}