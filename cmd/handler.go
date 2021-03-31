package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/mallvielfrass/fmc"
	"github.com/mallvielfrass/sessions/internal/crypto"
	uuid "github.com/satori/go.uuid"
)

func (st St) login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/registration/reg.html")
}
func (st St) Info(w http.ResponseWriter, r *http.Request) {
	data := "Info"
	tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
	tmpl.Execute(w, data)
}

type CreateUserJson struct {
	Success bool
	Login   string
	Session string
	Expiry  time.Time
	Error   string
}

func (st St) CreateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	obj := CreateUserJson{
		Success: false,
		Error:   "",
	}
	//r.ParseForm()
	fmt.Println("r: ", r.URL)
	login := r.URL.Query().Get("login")
	//login := r.FormValue("login")
	password := r.URL.Query().Get("password")
	fmc.Printfln("#bbt(CreateSession) #ybt Login:#bbt[#gbt%s#bbt] Password:#bbt[#gbt%s#bbt]", login, password)
	obj.Login = login
	user, err := st.base.GetUser(login)
	if err != nil {
		fmt.Printf("(CreateUserApi) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	m, err := crypto.ComparePassword(password, user.Hash)
	if err != nil {
		fmt.Printf("(CreateUserApi) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	if !m {
		fmt.Printf("(CreateUserApi) Error: %s \n", "password not compare")
		obj.Error = "password not compare"
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	fmt.Printf(" Login:%s |hash: %s \n", login, user.Hash)
	session := uuid.NewV4().String()
	myDate := time.Now()
	var month float64 = 3.0
	fmt.Println(myDate)
	shift := 2419200000000000 * time.Duration(month) //2419200000000000 = one month
	fmt.Println(shift)
	newDate := myDate.Add(shift)
	fmt.Println(newDate)
	err = st.base.CreateSession(login, session, int(newDate.Unix()))
	if err != nil {
		fmt.Printf("(CreateSession) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	fmt.Printf("UUIDv4: %s\n", session)
	obj.Success = true
	obj.Session = session
	obj.Expiry = newDate
	returnData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("(CreateSession) Error: %s \n", err)
		obj.Success = false
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	w.Write(returnData)
}
func (st St) CreateUserApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	obj := CreateUserJson{
		Success: false,
		Error:   "",
	}
	//r.ParseForm()
	fmt.Println("r: ", r.URL)
	login := r.URL.Query().Get("login")
	//login := r.FormValue("login")
	password := r.URL.Query().Get("password")
	fmc.Printfln("#bbt(CreateUserApi) #ybt Login:#bbt[#gbt%s#bbt] Password:#bbt[#gbt%s#bbt]", login, password)
	obj.Login = login

	hash, err := crypto.GeneratePassword(password)
	if err != nil {
		fmt.Printf("(CreateUserApi) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	err = st.base.CreateUser(login, hash)
	if err != nil {
		fmt.Printf("(CreateUserApi) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}

	fmt.Printf(" Login:%s |hash: %s \n", login, hash)
	session := uuid.NewV4().String()
	myDate := time.Now()
	var month float64 = 3.0
	fmt.Println(myDate)
	shift := 2419200000000000 * time.Duration(month) //2419200000000000 = one month
	fmt.Println(shift)
	newDate := myDate.Add(shift)
	fmt.Println(newDate)
	err = st.base.CreateSession(login, session, int(newDate.Unix()))
	if err != nil {
		fmt.Printf("(CreateSession) Error: %s \n", err)
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	fmt.Printf("UUIDv4: %s\n", session)
	obj.Success = true
	obj.Session = session
	obj.Expiry = newDate
	returnData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("(CreateSession) Error: %s \n", err)
		obj.Success = false
		obj.Error = err.Error()
		returnData, _ := json.Marshal(obj)
		w.Write(returnData)
		return
	}
	w.Write(returnData)
}
func (st St) CreateSessionApi(w http.ResponseWriter, r *http.Request) {
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
		fmt.Printf("(CreateSessionApi) Error: %s \n", err)
		fmt.Fprintf(w, "bad\n")
		return
	}

	fmt.Printf("UUIDv4: %s\n", session)

}
func (st St) SearchSessionApi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.FormValue("login")
	session := r.FormValue("session")
	uInfo, err := st.base.SearchSession(login, session)
	if err != nil {
		fmt.Printf("(SearchSessionApi) Error: %s \n", err)
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
