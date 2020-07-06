package pagecode

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

//HomePage function
func HomePage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "mainpage.html")
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
