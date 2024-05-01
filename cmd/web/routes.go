package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	mux.Get("/", app.Home)
	// mux.Post("/virtual-terminal-payment-succeeded", app.VirtualTerminalPaymentSucceeded)
	// mux.Get("/virtual-terminal-receipt", app.VirtualTerminalReceipt)

	mux.Get("/buy-dates/{id}", app.BuyOnce)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)
	mux.Get("/receipt", app.Receipt)

	mux.Get("/plans/bronze", app.BronzePlan)
	mux.Get("/receipt/bronze", app.BronzePlanReceipt)

	//Auhtentication 
	mux.Get("/signin", app.Signin)
	mux.Post("/signin", app.PostSignin)
	mux.Get("/signout", app.SignOut)
	
	//404 not found route
	mux.NotFound(app.PageNotFound)

	//Public file server
	publicFileServer := http.FileServer(http.Dir("./public/assets"))
	mux.Handle("/public/assets/*", http.StripPrefix("/public/assets", publicFileServer))


	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.Auth)
		mux.Get("/virtual-terminal", app.VirtualTerminal)
	})
	return mux
}
