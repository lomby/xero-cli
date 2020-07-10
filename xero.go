package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lomby/xero-cli/accounts"
	"github.com/lomby/xero-cli/auth"
	"github.com/lomby/xero-cli/xeroclient"
)

func main() {

	provider := auth.NewProvider()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	provider.GetAuthURL()

	r := chi.NewRouter()

	// r.HandleFunc("/", refreshToken)
	r.HandleFunc("/callback", provider.HandleCallback)
	r.HandleFunc("/tenants", xeroclient.GetTenants)
	r.HandleFunc("/invoice", accounts.GetInvoice)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err != nil {
		fmt.Println(err)
	}

	err = srv.ListenAndServe()

}
