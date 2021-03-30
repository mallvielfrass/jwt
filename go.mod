module github.com/mallvielfrass/sessions

go 1.16

//replace github.com/mallvielfrass/sessions/internal/handlers => ./internal/handlers

//replace github.com/mallvielfrass/sessions/internal/middleware => ./internal/middleware

require (
	github.com/go-chi/chi/v5 v5.0.2
	github.com/mallvielfrass/fmc v0.0.0-20210329150608-ca76733c5741
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 // indirect
	golang.org/x/term v0.0.0-20210317153231-de623e64d2a6 // indirect
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.6
)
