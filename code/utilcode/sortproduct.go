package utilcode

import "github.com/derpl-del/go-cloth/code/strcode"

// ByName sort
type ByName []strcode.ResponseProductList

func (data ByName) Len() int           { return len(data) }
func (data ByName) Less(i, j int) bool { return data[i].Productname < data[j].Productname }
func (data ByName) Swap(i, j int)      { data[i], data[j] = data[j], data[i] }

//ByDate sort
type ByDate []strcode.ResponseProductList

func (data ByDate) Len() int           { return len(data) }
func (data ByDate) Less(i, j int) bool { return data[i].Productupdate > data[j].Productupdate }
func (data ByDate) Swap(i, j int)      { data[i], data[j] = data[j], data[i] }
