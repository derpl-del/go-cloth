package ctrcode

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/go-cloth/code/pagecode"
	"github.com/derpl-del/go-cloth/code/sscode"
	"github.com/gorilla/mux"
)

//Funchandler func
func Funchandler() {
	fmt.Println("morning")
	sscode.CreateStore()
	r := mux.NewRouter()
	r.HandleFunc("/", pagecode.HomePage)
	r.HandleFunc("/login", pagecode.LoginPage)
	r.HandleFunc("/shopkeeper", pagecode.ShopKeeperPage)
	r.HandleFunc("/shopkeeper/product_add", pagecode.ProductAddPage)
	r.HandleFunc("/api/v1/preview", pagecode.PreviewImage)
	r.HandleFunc("/api/v1/productadd", pagecode.AddProductPage)
	r.HandleFunc("/SignInUser", pagecode.SignInHandler)
	r.HandleFunc("/SignUpUser", pagecode.SignUpHandler)
	r.HandleFunc("/GenerateToken", pagecode.GenerateToken).Methods("GET")
	r.HandleFunc("/api/v1/generateproductid", pagecode.GenerateProductID).Methods("GET")
	r.HandleFunc("/val/RemoveUserSession", pagecode.RemoveUserSession).Methods("GET")
	r.HandleFunc("/val/GetUserSession", pagecode.GetUserSession).Methods("GET")
	r.PathPrefix("/env/").Handler(http.StripPrefix("/env/", http.FileServer(http.Dir("env"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	r.PathPrefix("/data/").Handler(http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
