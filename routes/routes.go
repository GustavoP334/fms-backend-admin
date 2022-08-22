package routes

import (
	"application-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Decision)
	http.HandleFunc("/indexSemDocs", controllers.DecisionSemDocs)
	http.HandleFunc("/index", controllers.DecisionWithDocs)
	http.HandleFunc("/edit", controllers.Edit)
	// http.HandleFunc("/decision", controllers.Edit)
	//http.HandleFunc("/update", controllers.Update)
}
