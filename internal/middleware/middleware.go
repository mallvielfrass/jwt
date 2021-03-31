package middleware

import (
	"net/http"

	"github.com/mallvielfrass/fmc"
)

func Middleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmc.Printfln("#bbtmiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)
	next(rw, r)
}
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmc.Printfln("#bbtAuthMiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)

		next.ServeHTTP(w, r)
	})

}
func NoAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmc.Printfln("#bbtNoAuthMiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)

		next.ServeHTTP(w, r)
	})

}
