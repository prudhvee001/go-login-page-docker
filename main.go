package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl.Execute(w, nil)
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "mpr" && password == "Mpr001" {
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			tmpl.Execute(w, "Invalid credentials. Try again.")
		}
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, admin! You are successfully logged in."))
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
