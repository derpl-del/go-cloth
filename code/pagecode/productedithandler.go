package pagecode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/derpl-del/go-cloth/code/ddata"
	"github.com/derpl-del/go-cloth/code/jscode"
	"github.com/derpl-del/go-cloth/code/rdata"
	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/wdata"
)

//ReturnProductDetail for homepage
func ReturnProductDetail(w http.ResponseWriter, r *http.Request) {
	pid := r.FormValue("product_id")
	var article = rdata.GetProductData(pid)
	json.NewEncoder(w).Encode(article)
}

//DeleteProduct for homepage
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var resp strcode.ResponseForm
	pid := r.FormValue("product_id")
	fmt.Println(pid)
	err := ddata.DeleteProductData(pid)
	if err != nil {
		resp = strcode.ResponseForm{ErrorCode: "0001", ErrorMessage: fmt.Sprintf("%v", err)}
	} else {
		resp = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "success"}
	}
	json.NewEncoder(w).Encode(resp)
}

//EditProduct for homepage
func EditProduct(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	layoutFormat := "2006-01-02 15:04:05"
	date := currentTime.Format(layoutFormat)
	var ResStr strcode.ResponseForm
	ReqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(ReqBody))
	ReqStr := jscode.ProductRequest(ReqBody)
	article := rdata.GetProductData(ReqStr.Productcode)
	article.Productamount = ReqStr.Productamount
	article.Productname = ReqStr.Productname
	article.Productcategory = ReqStr.Productcategory
	article.Productgender = ReqStr.Productgender
	article.Productprice = "IDR " + ReqStr.Productprice
	article.Productsize = ReqStr.Productsize
	article.Productupdated = date
	err := wdata.EditProductData(article)
	if err != nil {
		ResStr = strcode.ResponseForm{ErrorCode: "9999", ErrorMessage: "Failed"}
	} else {
		ResStr = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "Success"}
	}
	json.NewEncoder(w).Encode(ResStr)
}
