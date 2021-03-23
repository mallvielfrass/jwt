package handlers

import (
	"fmt"
	"net/http"
)

type Res struct {
	R *http.Request
}

func (Res Res) Ck(cookie string) string {
	ckie, err := Res.R.Cookie(cookie)
	if err != nil {
		fmt.Println(err.Error())
		return "null"
	}
	return ckie.Value
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}
