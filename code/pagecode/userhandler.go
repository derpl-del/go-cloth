package pagecode

import (
	"encoding/json"
	"net/http"

	"github.com/derpl-del/go-cloth/code/sscode"
	"github.com/derpl-del/go-cloth/code/strcode"
)

//RemoveUserSession func
func RemoveUserSession(w http.ResponseWriter, r *http.Request) {
	sscode.ClearSession(w, r)
	ResStr := strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "success"}
	json.NewEncoder(w).Encode(ResStr)
}

//GetUserSession func
func GetUserSession(w http.ResponseWriter, r *http.Request) {
	user, token, role := sscode.GetUserName(r)
	ResStr := strcode.ResponseSession{User: user, ErrorCode: "0000", ErrorMessage: token, ErrorRole: role}
	json.NewEncoder(w).Encode(ResStr)
}
