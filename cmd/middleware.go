package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/mallvielfrass/fmc"
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
func Middleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmc.Printfln("#bbtmiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)
	next(rw, r)
}

func (st St) AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmc.Printfln("#bbtAuthMiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)
		login, loginErr := HandleCookie(r.Cookie("login"))
		if loginErr != nil {
			fmt.Println(loginErr)
		}
		session, sessionErr := HandleCookie(r.Cookie("session"))
		if sessionErr != nil {
			fmt.Println(sessionErr)
		}
		if sessionErr != nil || loginErr != nil {
			fmc.Printfln("#rbt(AuthMiddleware)> Error: #ybthandle cookie")
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		fmc.Printfln("Login: %s| Session: %s", login, session)
		inf, err := st.base.SearchSession(login, session)
		if err != nil {
			fmc.Printfln("#rbt(AuthMiddleware)> Error: #ybt%s", err)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		tm := int(time.Now().Unix())
		if inf.Expiry < tm {
			fmt.Printf("time now: %d\nexpir: %d\n", tm, inf.Expiry)
			fmt.Fprintf(w, "session expired\n")
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})

}
func (st St) NoAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmc.Printfln("#bbtNoAuthMiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)
		login, loginErr := HandleCookie(r.Cookie("login"))
		session, sessionErr := HandleCookie(r.Cookie("session"))
		if sessionErr == nil || loginErr == nil {
			fmc.Printfln("Login: %s| Session: %s", login, session)
			inf, err := st.base.SearchSession(login, session)
			if err == nil {
				tm := int(time.Now().Unix())
				if tm < inf.Expiry {
					http.Redirect(w, r, "/profile", http.StatusMovedPermanently)
				}
			}
			return
		}
		next.ServeHTTP(w, r)
	})

}
