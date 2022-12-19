package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/llRedXD/aluraRestApi/controllers"
	"github.com/llRedXD/aluraRestApi/middleware"
)

func HandleRequest()  {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/",controllers.Home)
	r.HandleFunc("/personalidades",controllers.Personalidade).Methods("Get")
	r.HandleFunc("/personalidades/{id}",controllers.UmaPersonalidade).Methods("Get")
	r.HandleFunc("/criaPersonalidade",controllers.CriaPersonalidade).Methods("Post")	
	r.HandleFunc("/deletePersonalidade/{id}",controllers.DeletePersonalidade).Methods("Delete")	
	r.HandleFunc("/editPersonalidade/{id}",controllers.EditPersonalidade).Methods("Put")	
	http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)) 

}