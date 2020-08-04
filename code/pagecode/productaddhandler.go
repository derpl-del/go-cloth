package pagecode

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/derpl-del/go-cloth/code/jscode"
	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/utilcode"
	"github.com/derpl-del/go-cloth/code/wdata"
)

//AddProductPage func
func AddProductPage(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	layoutFormat := "2006-01-02 15:04:05"
	date := currentTime.Format(layoutFormat)
	fmt.Println(date)
	var ResStr strcode.ResponseForm
	ReqBody, _ := ioutil.ReadAll(r.Body)
	ReqStr := jscode.ProductRequest(ReqBody)
	ReqStr.Productupdated = date
	ReqStr.Productcreated = date
	err := wdata.CreateProductData(ReqStr)
	if err != nil {
		ResStr = strcode.ResponseForm{ErrorCode: "9999", ErrorMessage: "Failed"}
	} else {
		ResStr = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "Success"}
	}
	json.NewEncoder(w).Encode(ResStr)
}

//PreviewImage func
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var token = r.FormValue("token")
	currentTime := time.Now()
	productid := currentTime.Format("200601021504-05")
	r.ParseMultipartForm(10 << 20)
	file, handler, err1 := r.FormFile("file")
	if handler.Filename != "" {
		if err1 != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err1)
			return
		}
		defer file.Close()
		path := "data/tmp/" + token
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}
		filetittle := path + "/" + productid + ".png"
		img, _ := os.Create(filetittle)
		defer img.Close()
		io.Copy(img, file)
		ResStr := strcode.ResponseUpload{ErrorCode: "0000", ErrorMessage: "", Location: filetittle, Productid: token}
		json.NewEncoder(w).Encode(ResStr)
	} else {
		ResStr := strcode.ResponseUpload{ErrorCode: "9999", ErrorMessage: "Failed", Location: ""}
		json.NewEncoder(w).Encode(ResStr)
	}
}

//GenerateProductID func
func GenerateProductID(w http.ResponseWriter, r *http.Request) {
	token1 := utilcode.GenerateString(1)
	currentTime := time.Now()
	productid := currentTime.Format("200601021504")
	tiket := strings.ToUpper(token1)
	token := tiket + "-" + productid
	ResStr := strcode.ResponseUpload{ErrorCode: "0000", ErrorMessage: "success", Location: token}
	json.NewEncoder(w).Encode(ResStr)
}
