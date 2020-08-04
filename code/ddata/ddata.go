package ddata

import (
	"fmt"
	"os"
)

//DeleteProductData func
func DeleteProductData(productid string) error {
	path := "product_list/" + productid
	fmt.Println(path)
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
