package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("./template/form.html"))
}

type studentSurvey struct {
	Name            string
	Email           string
	Age             string
	Role            string
	FavoriteFeature string
	Recommendation  string
	Comments        string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	ideas := r.Form["improvements"]
	recommendation := strings.Join(ideas, ", ")

	student := studentSurvey{
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Age:             r.FormValue("age"),
		Role:            r.FormValue("current-role"),
		FavoriteFeature: r.FormValue("radio-buttons"),
		Recommendation:  recommendation,
		Comments:        r.FormValue("comment"),
	}
	if r.FormValue("submit") == "Submit" {
		tmpl.Execute(w, struct {
			Success bool
			Message string
		}{Success: true, Message: "Survey Complete"})
	}
	fmt.Println(student)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}
