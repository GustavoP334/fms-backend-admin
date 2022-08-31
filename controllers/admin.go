package controllers

import (
	"application-web/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

/*var store = sessions.NewCookieStore([]byte("mysession"))

func Login(w http.ResponseWriter, r *http.Request) {
	allDrivers := models.GetAllDrivers()
	temp.ExecuteTemplate(w, "Login", allDrivers)
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == "abc" && password == "123" {
		session, _ := store.Get(r, "mysession")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		temp.Execute(w, data)
		fmt.Println(data)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}*/

func Index(w http.ResponseWriter, r *http.Request) {
	/*session, _ := store.Get(r, "mysession")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	fmt.Println(data)*/
	allDrivers := models.GetAllDrivers()
	temp.ExecuteTemplate(w, "Index", allDrivers)
}

func Decision(w http.ResponseWriter, r *http.Request) {
	oi := "oi"
	temp.ExecuteTemplate(w, "Decision", oi)
}

func DecisionSemDocs(w http.ResponseWriter, r *http.Request) {
	driverNoDoc := models.GetDriversNoDoc()
	temp.ExecuteTemplate(w, "IndexSemDocs", driverNoDoc)
}

func DecisionWithDocs(w http.ResponseWriter, r *http.Request) {
	driverWithDoc := models.GetDriversWithDoc()
	temp.ExecuteTemplate(w, "Index", driverWithDoc)
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var allRecords = make(map[string]interface{})
		for key, value := range r.Form {
			allRecords[key] = value
		}

		batata, err := json.Marshal(allRecords)
		if err != nil {
			fmt.Println("err", err.Error())
			return
		}

		batata = bytes.ReplaceAll(batata, []byte("["), []byte(""))
		batata = bytes.ReplaceAll(batata, []byte("]"), []byte(""))

		fmt.Println("teste", string(batata))
		models.CreateDocument(batata)
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var allRecords = make(map[string]interface{})
		for key, value := range r.Form {
			allRecords[key] = value
		}

		batata, err := json.Marshal(allRecords)
		if err != nil {
			fmt.Println("err", err.Error())
			return
		}
		userid := r.FormValue("driver_id")
		driverid := r.FormValue("driver_id")

		batata = bytes.ReplaceAll(batata, []byte("["), []byte(""))
		batata = bytes.ReplaceAll(batata, []byte("]"), []byte(""))

		fmt.Println("teste", string(batata))
		models.AtualizaDriver(batata, userid)
		models.AtualizaFieldsDocuments(batata, driverid)
	}
	http.Redirect(w, r, "/", 301)
}
