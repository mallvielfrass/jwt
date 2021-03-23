package middleware

import (
	"net/http"

	"github.com/mallvielfrass/fmc"
)

func Middleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmc.Printfln("#bbtmiddleware> #ybturl #bbt[#gbt%s#bbt] #ybtmethod #bbt[#gbt%s#bbt]", r.URL.Path, r.Method)
	next(rw, r)
	// do some stuff after
}
