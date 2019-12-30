package routes

import (
	"github.com/hammadtahirch/golang_basic_app/app/controllers"
	"github.com/hammadtahirch/golang_basic_app/app/middlewares"
	"github.com/gorilla/mux"
)

// AppRoute ...
type AppRoute struct {
	u controllers.UserController
	r controllers.RoleController
}

// Routes ... This function helps to maintain all the routes
func (a *AppRoute) Routes(r *mux.Router) {
	// simple web routes
	r.HandleFunc("/", controllers.ReactProduction)
	r.HandleFunc("/welcome", controllers.Welcome)
	//end

	// Api routes
	af := r.PathPrefix("/api/v1").Subrouter()
	af.Use(middlewares.LoggingMiddleware)
	af.HandleFunc("/signin", a.u.SignIn).Methods("POST")

	af.HandleFunc("/recover-password", a.u.RecoverPassword).Methods("POST")
	af.HandleFunc("/add-new-password", a.u.NewPassord).Methods("POST")
	af.HandleFunc("/registraion", a.u.Registration).Methods("POST")

	af.HandleFunc("/role-list", a.r.GetRoleList).Methods("GET")

	ar := r.PathPrefix("/api/v1").Subrouter()
	ar.Use(middlewares.ValidateJwtMiddlewear)
	ar.HandleFunc("/user", a.u.IndexHandler).Methods("GET")
	ar.HandleFunc("/user/{id:[0-9]+}", a.u.ShowHandler).Methods("GET")
	ar.HandleFunc("/user", a.u.StoreHandler).Methods("POST")
	ar.HandleFunc("/user/{id:[0-9]+}", a.u.UpdateHandler).Methods("PUT")
	ar.HandleFunc("/user/{id:[0-9]+}", a.u.DestroyHandler).Methods("DELETE")
	// end
}
