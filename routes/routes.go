package routes

import (
	"git-lab.boldapps.net/nifty-logix/mvc/app/controllers"
	"git-lab.boldapps.net/nifty-logix/mvc/app/middlewares"
	"github.com/gorilla/mux"
)

type AppRoute struct {
	u controllers.UserController
}

// Routes ... This function helps to maintain all the routes
func (a *AppRoute) Routes(r *mux.Router) {
	// simple web routes
	r.HandleFunc("/", controllers.Welcome)
	//end

	// Api routes
	s := r.PathPrefix("/api/v1").Subrouter()
	s.Use(middlewares.LoggingMiddleware)
	s.HandleFunc("/user", a.u.IndexHandler).Methods("GET")
	s.HandleFunc("/user/{id:[0-9]+}", a.u.ShowHandler).Methods("GET")
	s.HandleFunc("/user", a.u.StoreHandler).Methods("POST")
	s.HandleFunc("/user/{id:[0-9]+}", a.u.UpdateHandler).Methods("PUT")
	s.HandleFunc("/user/{id:[0-9]+}", a.u.DestroyHandler).Methods("DELETE")
	// end
}
