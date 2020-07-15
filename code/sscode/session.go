package sscode

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//StoreKey var
var StoreKey *sessions.CookieStore

//Sea const
const Sea = "SeassionPage"

//NewCookieStore Create Cookie
func NewCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	StoreKey := sessions.NewCookieStore(authKey, encryptionKey)
	StoreKey.Options.Path = "/"
	StoreKey.Options.MaxAge = 86400 * 7
	StoreKey.Options.HttpOnly = true

	return StoreKey
}

//SetSession for session
func SetSession(token string, username string, role bool, w http.ResponseWriter, r *http.Request) {
	session, _ := StoreKey.Get(r, Sea)
	session.Values["token"] = token
	session.Values["auth"] = true
	session.Values["user"] = username
	session.Values["role"] = role
	session.Save(r, w)
}

//GetUserName Session
func GetUserName(r *http.Request) (string, bool, bool) {
	session, _ := StoreKey.Get(r, Sea)
	if len(session.Values) == 0 {
		return "", false, false
	}
	return fmt.Sprintf("%s", session.Values["user"]), session.Values["auth"].(bool), session.Values["role"].(bool)
}

//ClearSession Session
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := StoreKey.Get(r, Sea)
	session.Options.MaxAge = -1
	session.Save(r, w)
}

//CreateStore Session
func CreateStore() {
	StoreKey = NewCookieStore()
}
