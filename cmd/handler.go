package main

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/mallvielfrass/sessions/internal/crypto"
	uuid "github.com/satori/go.uuid"
)

func HandleCookie(hCookie *http.Cookie, err error) (string, error) {
	if err != nil {
		return "", err
	}
	decodedValue, err := url.QueryUnescape(hCookie.Value)
	if err != nil {
		return "", err
	}
	return decodedValue, nil
}
func (st St) index(w http.ResponseWriter, r *http.Request) {
	user := st.base.GetAllUsers()
	msg := ""
	for i, m := range user {
		fmt.Printf("%d) user:%s , hash: %s\n", i, m.Login, m.Hash)
		msg = msg + fmt.Sprintf("%d) User: %s| Hash: %s\n", i, m.Login, m.Hash)
	}
	if len(user) == 0 {
		//	fmt.Fprintf(w, )
		//data := fmc.Sprint("#gbtUser not found")
		tmpl, _ := template.New("").Parse("<h1>User not found</h1>")
		tmpl.Execute(w, nil)
	} else {
		data := msg
		tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
		tmpl.Execute(w, data)
	}
}
func (st St) add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")

	hash, err := crypto.GeneratePassword(password)
	if err != nil {
		fmt.Printf("(add) Error: %s \n", err)
		fmt.Fprintf(w, "bad\n")
		return
	}
	err = st.base.CreateUser(login, hash)
	if err != nil {
		fmt.Printf("(add) Error: %s \n", err)
		fmt.Fprintf(w, "bad\n")
		return
	}

	fmt.Printf(" Login:%s |hash: %s \n", login, hash)
	fmt.Fprintf(w, "ok\n")
}
func (st St) t(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session := uuid.NewV4().String()

	login := r.FormValue("login")
	myDate := time.Now()
	var month float64 = 3.0
	fmt.Println(myDate)
	shift := 2419200000000000 * time.Duration(month) //2419200000000000 = one month
	fmt.Println(shift)
	newDate := myDate.Add(shift)
	fmt.Println(newDate)
	err := st.base.CreateSession(login, session, int(newDate.Unix()))
	if err != nil {
		fmt.Printf("(add) Error: %s \n", err)
		fmt.Fprintf(w, "bad\n")
		return
	}

	fmt.Printf("UUIDv4: %s\n", session)

}
func (st St) s(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//session := uuid.NewV4().String()

	login := r.FormValue("login")
	session := r.FormValue("session")
	//myDate := time.Now()
	//	var month float64 = 3.0
	//	fmt.Println(myDate)
	//shift := 2419200000000000 * time.Duration(month) //2419200000000000 = one month
	//	fmt.Println(shift)
	//	newDate := myDate.Add(shift)
	//	fmt.Println(newDate)
	uInfo, err := st.base.SearchSession(login, session)
	if err != nil {
		fmt.Printf("(add) Error: %s \n", err)
		fmt.Fprintf(w, "session not found\n")
		return
	}

	fmt.Printf("UUIDv4: %s\n", session)
	tm := int(time.Now().Unix())
	if uInfo.Expiry < tm {
		fmt.Printf("time now: %d\nexpir: %d\n", tm, uInfo.Expiry)
		fmt.Fprintf(w, "session expired\n")
		return
	}
	fmt.Fprintf(w, "ok\n")
}
