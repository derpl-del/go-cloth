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

//ResponseSession struct
type ResponseSession struct {
	User         string
	ErrorCode    string
	ErrorMessage bool
	ErrorRole    bool
}

//ResponseUpload struct
type ResponseUpload struct {
	Location     string
	Productid    string
	ErrorCode    string
	ErrorMessage string
}

//RequestProduct struct
type RequestProduct struct {
	Productname     string `json:"productname"`
	Producttoken    string `json:"Producttoken"`
	Productcode     string `json:"productcode"`
	Productprice    string `json:"productprice"`
	Productgender   string `json:"productgender"`
	Productsize     string `json:"productsize"`
	Productcategory string `json:"productcategory"`
	Productamount   string `json:"productamount"`
	Imagelocation   string `json:"imgsource"`
}
