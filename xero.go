package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/lomby/xero-cli/accounts"
	"github.com/lomby/xero-cli/auth"
	"github.com/lomby/xero-cli/xeroclient"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func appInfo() {
	app.Name = "Xero CLI"
	app.Usage = "Manage invoice, contacts and more..."
	app.Version = "1.0.0"
}

func main() {

	provider := auth.NewProvider()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// provider.GetAuthURL()

	appInfo()

	var contactID string
	var invoiceID string

	contactIDFlag := &cli.StringFlag{
		Name:        "contactId",
		Usage:       "Sets the contact id",
		Destination: &contactID,
		Required:    true,
	}
	invoiceIDFlag := &cli.StringFlag{
		Name:        "invoiceId",
		Usage:       "Sets the invoice id",
		Destination: &invoiceID,
		Required:    true,
	}

	// Commands and Subcommands

	app.Commands = []*cli.Command{
		// Commands for Customers in Google Reseller API
		&cli.Command{
			Name:        "invoice",
			Usage:       "Invoice commands for Xero API",
			Description: "Various invoice commands over Xero API",
			Subcommands: []*cli.Command{
				{
					Name:        "single",
					Usage:       "invoice single --invoiceId *****",
					Description: "retrieve a single invoice by providing an invoice ID",
					Category:    "invoice",
					Flags:       []cli.Flag{invoiceIDFlag},
					Action: func(c *cli.Context) error {
						invoice, err := accounts.GetInvoice(invoiceID)
						if err != nil {
							return err
						}
						fmt.Println(invoice)
						return nil
					},
				},
				{
					Name:        "all",
					Usage:       "invoice all --contactId *****",
					Description: "retrieve all invoices by providing an contact ID",
					Category:    "invoice",
					Flags:       []cli.Flag{contactIDFlag},
					Action: func(c *cli.Context) error {
						invoices, err := accounts.GetInvoices(contactID)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(invoices)
						return nil
					},
				},
				{
					Name:        "link",
					Usage:       "invoice link --invoiceId *****",
					Description: "retrieve the online invoice link",
					Category:    "invoice",
					Flags:       []cli.Flag{invoiceIDFlag},
					Action: func(c *cli.Context) error {
						link, err := accounts.GetInvoiceLink(invoiceID)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(link)
						return nil
					},
				},
			},
		},
	}

	r := chi.NewRouter()

	// r.HandleFunc("/", refreshToken)
	r.HandleFunc("/callback", provider.HandleCallback)
	r.HandleFunc("/tenants", xeroclient.GetTenants)

	// srv := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      r,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// 	IdleTimeout:  120 * time.Second,
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = srv.ListenAndServe()

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
