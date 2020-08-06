package pagecode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/derpl-del/go-cloth/code/rdata"
	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/utilcode"
)

//ReturnAllProduct for homepage
func ReturnAllProduct(w http.ResponseWriter, r *http.Request) {
	var ListProduct []strcode.ResponseProductList
	var data strcode.ResponseProductList
	page := r.FormValue("page")
	pageint, _ := strconv.Atoi(page)
	itemperpage := 9
	pagen := ((pageint * itemperpage) + 1)
	pageb := (((pageint - 1) * itemperpage) + 1)
	fileInfo, err := ioutil.ReadDir("product_list/")
	if err != nil {
		fmt.Println(err)
	}
	for i, info := range fileInfo {
		num := i + 1
		if num >= int(pageb) && num < int(pagen) {
			var article = rdata.GetProductData(info.Name())
			data.Productnum = num
			data.Productcode = article.Producttoken
			data.Productname = article.Productname
			data.Imagelocation = article.Imagelocation
			data.Productupdate = article.Productupdated
			ListProduct = append(ListProduct, data)
		}
	}
	lenfile := float64(len(fileInfo))
	lenpage := float64(itemperpage)
	totalpage := math.Ceil(lenfile / lenpage)
	sort.Sort(utilcode.ByDate(ListProduct))
	Articles := strcode.AllResponseProduct{ListProduct: ListProduct, Page: pageint, Totalpage: int(totalpage)}
	json.NewEncoder(w).Encode(Articles)
}

//ReturnAllsProduct for homepage
func ReturnAllsProduct(w http.ResponseWriter, r *http.Request) {
	var ListProduct []strcode.ResponseProductList
	var data strcode.ResponseProductList
	fileInfo, err := ioutil.ReadDir("product_list/")
	if err != nil {
		fmt.Println(err)
	}
	for i, info := range fileInfo {
		var article = rdata.GetProductData(info.Name())
		data.Productnum = i + 1
		data.Productcode = article.Producttoken
		data.Productname = article.Productname
		data.Imagelocation = article.Imagelocation
		data.Productupdate = article.Productupdated
		ListProduct = append(ListProduct, data)
	}
	sort.Sort(utilcode.ByDate(ListProduct))
	Articles := strcode.AllResponseProduct{ListProduct: ListProduct}
	json.NewEncoder(w).Encode(Articles)
}
