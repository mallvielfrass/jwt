package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mallvielfrass/fmc"
	"github.com/mallvielfrass/jwt/internal/handlers"
	"github.com/mallvielfrass/jwt/internal/middleware"
	"github.com/urfave/negroni"
)

var (
	port = "8080"
)

func inx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<head></head>")
	//v := Res{R: r}
	//tcookie := v.ck("appointment")
	for _, cookie := range r.Cookies() {
		fmt.Fprintf(w, "<br>Found a cookie named: name:%s, value: %s", cookie.Name, cookie.Value)
	}
	//text := "<p>some text</p>"
	fmt.Fprintf(w, "<p><a href='dog.html'>Собаки</a></p>")

	//tt := `{{.}}`

	//t := template.Must(template.New("test").Parse(tt))
	//t.Execute(w, text)
	w.Header().Set("Content-Type", "text/html")
}
func main() {
	r := mux.NewRouter()
	//handlers
	r.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	r.HandleFunc("/", inx).Methods("GET")

	fmc.Printfln("#gbtApp running on [#ybt:%s#gbt] port.", port)

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middleware.Middleware))
	n.UseHandler(r)
	http.ListenAndServe(":"+port, n)
}
