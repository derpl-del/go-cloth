package pagecode

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/derpl-del/go-cloth/code/dbcode"
	"github.com/derpl-del/go-cloth/code/jscode"
	"github.com/derpl-del/go-cloth/code/sscode"
	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/utilcode"
)

//SignInHandler func
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var ResStr strcode.ResponseForm
	ReqBody, _ := ioutil.ReadAll(r.Body)
	ReqStr := jscode.UsernameInfo(ReqBody)
	UserVal := dbcode.SelectUserData("username", "username", ReqStr.Username)
	Role := dbcode.SelectUserData("role", "username", ReqStr.Username)
	if UserVal == false {
		ResStr = strcode.ResponseForm{ErrorCode: "0001", ErrorMessage: "Invalid Username"}
	} else {
		PassVal := dbcode.SelectUserValueData("userpassword", "username", ReqStr.Username, "userpassword", ReqStr.Password)
		if PassVal == false {
			ResStr = strcode.ResponseForm{ErrorCode: "0002", ErrorMessage: "Invalid Password"}
		} else {
			ResStr = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "Success"}
			sscode.SetSession(ReqStr.Ticket, ReqStr.Username, Role, w, r)
		}
	}
	json.NewEncoder(w).Encode(ResStr)
}

//SignUpHandler func
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var ResStr strcode.ResponseForm
	ReqBody, _ := ioutil.ReadAll(r.Body)
	ReqStr := jscode.UsernameInfo(ReqBody)
	UserVal := dbcode.SelectUserData("username", "username", ReqStr.Username)
	if UserVal == false && len(ReqStr.Password) >= 5 && len(ReqStr.Username) >= 5 {
		ResStr = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "success"}
		dbcode.InsertUserData(ReqStr.Username, ReqStr.Password)
	} else if UserVal == true {
		ResStr = strcode.ResponseForm{ErrorCode: "1001", ErrorMessage: "username already exists"}
	} else if len(ReqStr.Username) <= 5 {
		ResStr = strcode.ResponseForm{ErrorCode: "1003", ErrorMessage: "username lenght min 5"}
	} else if len(ReqStr.Password) <= 5 {
		ResStr = strcode.ResponseForm{ErrorCode: "1004", ErrorMessage: "password lenght min 5"}
	} else {
		ResStr = strcode.ResponseForm{ErrorCode: "1005", ErrorMessage: "Internal Error"}
	}
	json.NewEncoder(w).Encode(ResStr)
}

//GenerateToken func
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	token1 := utilcode.GenerateString(6)
	token2 := utilcode.GenerateString(6)
	token3 := utilcode.GenerateString(6)
	token4 := utilcode.GenerateString(6)
	token5 := utilcode.GenerateString(6)
	token := token1 + "-" + token2 + "-" + token3 + "-" + token4 + "-" + token5
	ResStr := strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: token}
	json.NewEncoder(w).Encode(ResStr)
}
