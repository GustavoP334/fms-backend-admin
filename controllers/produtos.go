package controllers

import (
	"application-web/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allDrivers := models.GetAllDrivers()
	temp.ExecuteTemplate(w, "Index", allDrivers)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	driverId := r.URL.Query().Get("id")
	driver := models.EditDriver(driverId)
	verify := driver.FieldsDocuments.DriverID
	if len(verify) == 0 {
		temp.ExecuteTemplate(w, "Create", driver)
	} else {
		temp.ExecuteTemplate(w, "Edit", driver)
	}

}

/*func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
*/
