package main

import (
	"log"
	"net/http"

	"OAuth/app"
	"OAuth/routes/callback"
	"OAuth/routes/home"
	"OAuth/routes/login"
	"OAuth/routes/logout"
	middlewares "OAuth/routes/middleware"
	"OAuth/routes/user"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {

	err := app.Init()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler).Methods("GET")
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe(":3000", r))
}
