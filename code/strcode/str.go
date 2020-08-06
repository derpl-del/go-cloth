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
	Productnum      int    `json:"productnum"`
	Productname     string `json:"productname"`
	Producttoken    string `json:"Producttoken"`
	Productcode     string `json:"productcode"`
	Productprice    string `json:"productprice"`
	Productgender   string `json:"productgender"`
	Productsize     string `json:"productsize"`
	Productcategory string `json:"productcategory"`
	Productamount   string `json:"productamount"`
	Imagelocation   string `json:"imgsource"`
	Productcreated  string `json:"productcreated"`
	Productupdated  string `json:"Productupdated"`
}

//AllResponseProduct struct
type AllResponseProduct struct {
	ListProduct []ResponseProductList
	Page        int
	Totalpage   int
}

//ResponseProductList struct
type ResponseProductList struct {
	Productnum    int    `json:"productnum"`
	Productname   string `json:"productname"`
	Productcode   string `json:"productcode"`
	Productupdate string `json:"productupdate"`
	Imagelocation string `json:"imgsource"`
}

//RequestHighLight struct
type RequestHighLight struct {
	ListProduct []ProductHL `json:"ListProduct"`
}

//ProductHL struct
type ProductHL struct {
	Productcode string `json:"productcode"`
	Imgloc      string `json:"imgloc"`
}
