package routes

import (
	"application-web/controllers"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"os"
)

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		validUsername := os.Getenv("USUARIO")
		validPassword := os.Getenv("SENHA")
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(validUsername))
			expectedPasswordHash := sha256.Sum256([]byte(validPassword))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func CarregaRotas() {
	http.HandleFunc("/", basicAuth(controllers.Decision))
	http.HandleFunc("/indexSemDocs", basicAuth(controllers.DecisionSemDocs))
	http.HandleFunc("/index", basicAuth(controllers.DecisionWithDocs))
	http.HandleFunc("/edit", basicAuth(controllers.Edit))
	http.HandleFunc("/insert", basicAuth(controllers.Insert))
	//http.HandleFunc("/decision", controllers.Edit)
	//http.HandleFunc("/update", controllers.Update)
}
