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
