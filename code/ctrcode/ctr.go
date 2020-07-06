package ctrcode

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/go-cloth/code/pagecode"
	"github.com/gorilla/mux"
)

//Funchandler func
func Funchandler() {
	fmt.Println("morning")
	r := mux.NewRouter()
	r.HandleFunc("/", pagecode.HomePage)
	r.PathPrefix("/env/").Handler(http.StripPrefix("/env/", http.FileServer(http.Dir("env"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
