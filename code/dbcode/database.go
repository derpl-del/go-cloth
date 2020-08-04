package dbcode

import (
	"database/sql"
	"fmt"

	//for framework
	"github.com/derpl-del/go-cloth/code/strcode"
	//for framework
	_ "github.com/godror/godror"
)

//OpenConnection func
func OpenConnection() *sql.DB {
	db, err := sql.Open("godror", "REINDEER/banana1@xe")
	if err != nil {
		fmt.Println(err)
		return db
	}
	return db
}

//InsertUserData func
func InsertUserData(UserName string, UserPassword string) {
	db := OpenConnection()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO USER_INFO (USERNAME,USERPASSWORD,CREATED_DATE,ROLE) VALUES ('%s', '%s',sysdate,'customer')", UserName, UserPassword)
	//fmt.Println(statementSQL)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}

//SelectUserData func
func SelectUserData(Coloumn string, Validation string, Data string) bool {
	db := OpenConnection()
	defer db.Close()
	statementSQL := fmt.Sprintf("SELECT %s FROM USER_INFO WHERE %s = '%s'", Coloumn, Validation, Data)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return false
	}
	var Rs1 string
	for rows.Next() {
		rows.Scan(&Rs1)
	}
	if Rs1 != "" {
		return true
	}
	return false
}

//SelectUserValueData func
func SelectUserValueData(Coloumn string, Validation1 string, Data1 string, Validation2 string, Data2 string) bool {
	db := OpenConnection()
	defer db.Close()
	statementSQL := fmt.Sprintf("SELECT %s FROM USER_INFO WHERE %s = '%s' and %s = '%s'", Coloumn, Validation1, Data1, Validation2, Data2)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return false
	}
	var Rs1 string
	for rows.Next() {
		rows.Scan(&Rs1)
	}
	if Rs1 != "" {
		return true
	}
	return false
}

//InserProductData func
func InserProductData(Input strcode.RequestProduct) {
	db := OpenConnection()
	defer db.Close()
	statementSQL := fmt.Sprintf("INSERT INTO PRODUCT_INFO (PRODUCT_ID, PRODUCT_NAME, PRODUCT_CATEGORY, PRODUCT_SIZE, PRODUCT_GENDER, AMOUNT, CREATED_DATE, STATUS) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', SYSDATE, '%v')", Input.Productcode, Input.Productname, Input.Productcategory, Input.Productsize, Input.Productgender, Input.Productamount, "ACTIVE")
	fmt.Println(statementSQL)
	rows, err := db.Query(statementSQL)
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
	}
	defer rows.Close()
}
