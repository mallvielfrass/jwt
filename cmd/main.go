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
	r.HandleFunc("/login", st.login)
	r.Route("/api", func(r chi.Router) {
		//only to this method must be acces for not authorized user
		r.With(middleware.NoAuthMiddleware).Route("/", func(r chi.Router) {
			r.HandleFunc("/create_user", st.CreateUserApi)
			r.HandleFunc("/create_session", st.CreateSessionApi)
			r.HandleFunc("/search_session", st.SearchSessionApi) //for tesing
		})

		//to this methods must be acces for only authorized user
		r.With(middleware.AuthMiddleware).Route("/", func(r chi.Router) {
			r.HandleFunc("/info", st.Info)
		})
	})
	//logger
	n := negroni.New()
	n.Use(negroni.HandlerFunc(middleware.Middleware))
	n.UseHandler(r)
	fmc.Printfln("#gbtApp running on [#ybt:%s#gbt] port.", port)
	http.ListenAndServe(":"+port, n)
}
