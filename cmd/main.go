package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mallvielfrass/fmc"
	"github.com/mallvielfrass/sessions/internal/database"
	"github.com/mallvielfrass/sessions/internal/middleware"
	"github.com/urfave/negroni"
)

type St struct {
	base database.Database
}

func main() {
	port := "8080"
	r := chi.NewRouter()
	base := database.Init("base.db")
	base.CreateUserTable()
	base.CreateSessionTable()
	st := St{
		base: base,
	}
	//handlers
	r.HandleFunc("/", st.index)
	r.HandleFunc("/create_user", st.add)
	r.HandleFunc("/t", st.t)
	r.HandleFunc("/s", st.s)
	n := negroni.New()
	n.Use(negroni.HandlerFunc(middleware.Middleware))
	n.UseHandler(r)
	fmc.Printfln("#gbtApp running on [#ybt:%s#gbt] port.", port)
	http.ListenAndServe(":"+port, n)
}
