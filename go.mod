module github.com/mallvielfrass/sessions

go 1.16

replace github.com/mallvielfrass/sessions/internal/handlers => ./internal/handlers

replace github.com/mallvielfrass/sessions/internal/middleware => ./internal/middleware

require (
	github.com/gorilla/mux v1.8.0
	github.com/mallvielfrass/fmc v0.0.0-20210319211811-9067867fc527
	github.com/mallvielfrass/sessions/internal/handlers v0.0.0-00010101000000-000000000000
	github.com/mallvielfrass/sessions/internal/middleware v0.0.0-00010101000000-000000000000
	github.com/urfave/negroni v1.0.0
)
