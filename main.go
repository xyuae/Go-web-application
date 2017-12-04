package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		p := Page{Name: "Tony"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})	// http.HandleFunc

	fmt.Println(http.ListenAndServe(":8080", nil))
}
