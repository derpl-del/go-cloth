package pagecode

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

//ProductAddPage function
func ProductAddPage(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("view", "productaddpage.html")
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
