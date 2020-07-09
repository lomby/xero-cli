package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lomby/xero-cli/auth"
	"golang.org/x/oauth2"
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
	r.HandleFunc("/tenants", getTenants)
	r.HandleFunc("/invoice", getInvoice)

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

func getTenants(res http.ResponseWriter, req *http.Request) {

	provider := auth.NewProvider()

	js, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	var credentials oauth2.Token
	json.Unmarshal(js, &credentials)

	client := provider.Config.Client(provider.Ctx, &credentials)

	resp, err := client.Get("https://api.xero.com/connections")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

}

func getInvoice(res http.ResponseWriter, req *http.Request) {

	invoiceID := req.FormValue("id")

	provider := auth.NewProvider()

	js, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	var credentials oauth2.Token
	json.Unmarshal(js, &credentials)

	client := provider.Config.Client(provider.Ctx, &credentials)

	request, err := http.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/" + invoiceID, nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("xero-tenant-id", os.Getenv("UK_TENANT_ID"))

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

}
