package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mallvielfrass/fmc"
	"github.com/mallvielfrass/sessions/internal/crypto"
	"github.com/mallvielfrass/sessions/internal/database"

	"github.com/urfave/negroni"
)

type St struct {
	base database.Database
}

func (st St) index(w http.ResponseWriter, r *http.Request) {
	login, loginErr := HandleCookie(r.Cookie("login"))
	if loginErr != nil {
		fmt.Println(loginErr)
	}
	session, sessionErr := HandleCookie(r.Cookie("session"))
	if sessionErr != nil {
		fmt.Println(sessionErr)
	}
	if sessionErr != nil || loginErr != nil {
		fmc.Printfln("#rbt(index)> Error: #ybthandle cookie")
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}
	fmc.Printfln("Login: %s| Session: %s", login, session)
	fmt.Fprint(w, "index")
}

func staticRouter(w http.ResponseWriter, r *http.Request) {
	urlFile := r.URL.Path
	ext := crypto.GetType(urlFile)
	info, area := crypto.CheckAccessArea("." + urlFile)
	if !area {
		fmc.Printfln("#bbtStaticRouter> #rbtError: #ybtURL not in access area:#bbt[#gbt%s#bbt]", urlFile)
		http.NotFound(w, r)
		return
	}
	fmc.Printfln("#bbtStaticRouter> #ybtURL:#bbt[#gbt%s#bbt] #ybtType:#bbt[#gbt%s#bbt] #ybtLocal File:#bbt[#gbt%s#bbt]", urlFile, ext, info)
	switch ext {
	case "css":
		w.Header().Set("Content-Type", "text/css")
		fmc.Printfln("#bbtStaticRouter> #ybtSet header: #bbt[#gbt%s#bbt]", ext)
	case "js":
		w.Header().Set("Content-Type", "application/javascript")
		fmc.Printfln("#bbtStaticRouter> #ybtSet header: #bbt[#gbt%s#bbt]", ext)
	case "ttf":
		w.Header().Set("Content-Type", "application/x-font-ttf")
		fmc.Printfln("#bbtStaticRouter> #ybtSet header: #bbt[#gbt%s#bbt]", ext)
	default:
		fmc.Printfln("#bbtStaticRouter> Undefined type [%s] of file: [%s]", ext, urlFile)
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, info)
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
	r.HandleFunc("/web/*", staticRouter)
	r.With(st.AuthMiddleware).HandleFunc("/profile", st.Info)
	r.Route("/api", func(r chi.Router) {
		//only to this method must be acces for not authorized user
		r.With(st.NoAuthMiddleware).Route("/auth", func(r chi.Router) {
			r.HandleFunc("/create_user", st.CreateUserApi)
			r.HandleFunc("/create_session", st.CreateSession)
			//r.HandleFunc("/search_session", st.SearchSessionApi) //for tesing
		})

		//to this methods must be acces for only authorized user
		r.With(st.AuthMiddleware).Route("/access", func(r chi.Router) {
			r.HandleFunc("/info", st.Info)
		})
	})
	//logger
	n := negroni.New()
	n.Use(negroni.HandlerFunc(Middleware))
	n.UseHandler(r)
	fmc.Printfln("#gbtApp running on [#ybt:%s#gbt] port.", port)
	http.ListenAndServe(":"+port, n)
}
