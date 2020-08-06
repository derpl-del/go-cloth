package pagecode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/derpl-del/go-cloth/code/jscode"
	"github.com/derpl-del/go-cloth/code/rdata"
	"github.com/derpl-del/go-cloth/code/strcode"
	"github.com/derpl-del/go-cloth/code/wdata"
)

//HighLightCreate func
func HighLightCreate(w http.ResponseWriter, r *http.Request) {
	var ResStr strcode.ResponseForm
	ReqBody, _ := ioutil.ReadAll(r.Body)
	ReqStr := jscode.HighLighRequest(ReqBody)
	if ReqStr.ListProduct == nil {
		ResStr = strcode.ResponseForm{ErrorCode: "9999", ErrorMessage: "Invalid Request"}
	} else {
		err := wdata.CreateHL(ReqStr)
		if err != nil {
			ResStr = strcode.ResponseForm{ErrorCode: "9999", ErrorMessage: fmt.Sprintf("%v", err)}
		}
		ResStr = strcode.ResponseForm{ErrorCode: "0000", ErrorMessage: "success"}
	}
	json.NewEncoder(w).Encode(ResStr)
}

//HighLightRead func
func HighLightRead(w http.ResponseWriter, r *http.Request) {
	res := rdata.GetHLData()
	json.NewEncoder(w).Encode(res)
}
