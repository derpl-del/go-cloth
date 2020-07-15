package wdata

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/utilcode"
)

//CreateProductData func
func CreateProductData(input strcode.RequestProduct) error {
	path := "product_list/" + input.Producttoken
	productcode := utilcode.SubsAfer(input.Productcode, "Product Code : ")
	input.Productcode = productcode
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
	pathjs := path + "/js"
	if _, err := os.Stat(pathjs); os.IsNotExist(err) {
		os.Mkdir(pathjs, 0700)
	}
	pathimg := path + "/img"
	if _, err := os.Stat(pathimg); os.IsNotExist(err) {
		os.Mkdir(pathimg, 0700)
	}
	location := utilcode.SubsAfer(input.Imagelocation, "../")
	oldLocation := location
	newLocation := pathimg + "/" + productcode + ".png"
	input.Imagelocation = newLocation
	err2 := os.Rename(oldLocation, newLocation)
	if err2 != nil {
		return nil
	}
	JSTittle := pathjs + "/product_" + productcode + ".json"
	out, _ := json.Marshal(input)
	err := ioutil.WriteFile(JSTittle, out, 0777)
	if err != nil {
		return err
	}
	removeloc := "data/tmp/" + productcode
	err3 := os.Remove(removeloc)
	if err3 != nil {
		return err3
	}
	return nil
}
