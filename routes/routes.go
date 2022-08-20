package routes

import (
	"application-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/edit", controllers.Edit)
	//http.HandleFunc("/update", controllers.Update)
}
