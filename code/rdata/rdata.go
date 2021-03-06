package rdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/derpl-del/go-cloth/code/strcode"
)

//GetProductData JSON
func GetProductData(input string) strcode.RequestProduct {
	JSTittle := "product_list/" + input + "/js/" + input + ".json"
	jsonFile, err := os.Open(JSTittle)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var struct1 strcode.RequestProduct
	json.Unmarshal(byteValue, &struct1)
	return struct1
}

//GetHLData JSON
func GetHLData() strcode.RequestHighLight {
	path := "img/HL/HL_List.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var struct1 strcode.RequestHighLight
	json.Unmarshal(byteValue, &struct1)
	return struct1
}
