package pagecode

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

//ProductEditPage function
func ProductEditPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "producteditpage.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
	err = tmpl.Execute(w, "test")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
}
