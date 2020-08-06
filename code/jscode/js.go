package jscode

import (
	"encoding/json"

	"github.com/derpl-del/go-cloth/code/strcode"
)

//UsernameInfo JSON
func UsernameInfo(input []byte) strcode.UsernameInfo {
	var response strcode.UsernameInfo
	json.Unmarshal(input, &response)
	return response
}

//ProductRequest JSON
func ProductRequest(input []byte) strcode.RequestProduct {
	var response strcode.RequestProduct
	json.Unmarshal(input, &response)
	return response
}

//HighLighRequest JSON
func HighLighRequest(input []byte) strcode.RequestHighLight {
	var response strcode.RequestHighLight
	json.Unmarshal(input, &response)
	return response
}
