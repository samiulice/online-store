package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Post("/api/payment-intent", app.GetPaymentIntent)
	mux.Post("/api/create-customer-and-subscribe-to-plan", app.CreateCustomerAndSubscribe)
	mux.Post("/api/authenticate", app.CreateAuthToken)
	mux.Post("/api/is-authenticated", app.CheckAuthenticated)
	mux.Post("/api/forgot-password", app.ForgotPassword)
	mux.Post("/api/reset-password", app.ResetPassword)

	//Secure routes
	mux.Route("/api/admin", func(mux chi.Router) {
		mux.Use(app.Auth)

		mux.Post("/virtual-terminal-payment-succeeded", app.VirtualTerminalPaymentSucceeded)
		
		//Order
		mux.Post("/analytics/order/view/{type}", app.GetOrdersHistoy)
		mux.Post("/analytics/order/refund/{type}", app.RefundCharge)
		
		//Order
		mux.Post("/analytics/subscription/cancel/{type}", app.CancelSubscription)
		
		mux.Post("/analytics/transaction/view/{type}", app.GetTransactionHistory)

		//cusotmer
		mux.Post("/customer/profile/view/{id}", app.AdminCustomerProfile)
	})
	return mux
}
