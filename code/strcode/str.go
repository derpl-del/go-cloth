package strcode

//UsernameInfo struct
type UsernameInfo struct {
	Ticket   string `json:"ticket"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//ResponseForm struct
type ResponseForm struct {
	ErrorCode    string
	ErrorMessage string
}
