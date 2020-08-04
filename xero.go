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
	var contactData string
	var invoiceID string
	var invoiceData string
	var paymentData string
	var creditNoteData string
	var pdf bool
	var search string

	contactIDFlag := &cli.StringFlag{
		Name:        "contactId",
		Usage:       "Sets the contact id",
		Destination: &contactID,
		Required:    true,
	}
	contactDataFlag := &cli.StringFlag{
		Name:        "contactData",
		Usage:       "Contact json data",
		Destination: &contactData,
		Required:    true,
	}
	invoiceIDFlag := &cli.StringFlag{
		Name:        "invoiceId",
		Usage:       "Sets the invoice id",
		Destination: &invoiceID,
		Required:    true,
	}
	invoiceDataFlag := &cli.StringFlag{
		Name:        "invoiceData",
		Usage:       "Invoice Json Data",
		Destination: &invoiceData,
		Required:    true,
	}
	creditNoteDataFlag := &cli.StringFlag{
		Name:        "paymentData",
		Usage:       "Payment Json Data",
		Destination: &paymentData,
		Required:    true,
	}
	paymentDataFlag := &cli.StringFlag{
		Name:        "creditNoteData",
		Usage:       "Credit Note Json Data",
		Destination: &creditNoteData,
		Required:    true,
	}
	pdfFlag := &cli.BoolFlag{
		Name:        "pdf",
		Usage:       "Invoice Json Data",
		Destination: &pdf,
		Required:    false,
	}
	searchFlag := &cli.StringFlag{
		Name:        "search",
		Usage:       "Search Term (e.g. Name=John Doe or )",
		Destination: &search,
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
					Flags:       []cli.Flag{invoiceIDFlag, pdfFlag},
					Action: func(c *cli.Context) error {
						invoice, err := accounts.GetInvoice(invoiceID, pdf)
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
							return nil
						}
						fmt.Println(invoices)
						return nil
					},
				},
				{
					Name:        "create",
					Usage:       "invoice create --invoiceData {}",
					Description: "Creates a new invoice using invoice json data",
					Category:    "invoice",
					Flags:       []cli.Flag{invoiceDataFlag},
					Action: func(c *cli.Context) error {
						invoice, err := accounts.CreateInvoice(invoiceData)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(invoice)
						return nil
					},
				},
				{
					Name:        "update",
					Usage:       "invoice update --invoiceData {}",
					Description: "Updates an existing invoice using invoice json data",
					Category:    "invoice",
					Flags:       []cli.Flag{invoiceDataFlag},
					Action: func(c *cli.Context) error {
						invoice, err := accounts.CreateInvoice(invoiceData)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(invoice)
						return nil
					},
				},
				{	paymentDataFlag := &cli.StringFlag{
		Name:        "paymentData",
		Usage:       "Payment Json Data",
		Destination: &paymentData,
		Required:    true,
	}
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
		&cli.Command{
			Name:        "contact",
			Usage:       "Contact commands for Xero API",
			Description: "Various contact commands over Xero API",
			Subcommands: []*cli.Command{
				{
					Name:        "get",
					Usage:       "contact get --contactId *****",
					Description: "retrieve a single Contact by providing a contact ID",
					Category:    "contact",
					Flags:       []cli.Flag{contactIDFlag},
					Action: func(c *cli.Context) error {
						contact, err := accounts.GetContact(contactID, "")
						if err != nil {
							return err
						}
						fmt.Println(contact)
						return nil
					},
				},
				{
					Name:        "search",
					Usage:       "contact search --search (e.g {\"Name\": \"ABC Trading\"} )",
					Description: "search a single Contact by providing a search key and value",
					Category:    "contact",
					Flags:       []cli.Flag{searchFlag},
					Action: func(c *cli.Context) error {
						contact, err := accounts.GetContact("", search)
						if err != nil {
							return err
						}
						fmt.Println(contact)
						return nil
					},
				},
				{
					Name:        "update",
					Usage:       "contact update --contactData {}",
					Description: "Updates contact information by providing contact json data",
					Category:    "contact",
					Flags:       []cli.Flag{contactDataFlag},
					Action: func(c *cli.Context) error {
						contact, err := accounts.GetContact("", contactData)
						if err != nil {
							return err
						}
						fmt.Println(contact)
						return nil
					},
				},
			},
		},
		&cli.Command{
			Name:        "payment",
			Usage:       "Payment commands for Xero API",
			Description: "Various payment commands over Xero API",
			Subcommands: []*cli.Command{
				{
					Name:        "create",
					Usage:       "payment create --paymentData {}",
					Description: "Create a payment by providing payment data",
					Category:    "payment",
					Flags:       []cli.Flag{paymentDataFlag},
					Action: func(c *cli.Context) error {
						contact, err := accounts.MakePayment(paymentData)
						if err != nil {
							return err
						}
						fmt.Println(contact)
						return nil
					},
				},
			},
		},
		&cli.Command{
			Name:        "creditnote",
			Usage:       "Credit Note commands for Xero API",
			Description: "Various credit note commands over Xero API",
			Subcommands: []*cli.Command{
				{
					Name:        "create",
					Usage:       "creditnote create --creditNoteData {}",
					Description: "Create a credit note by providing Credit note data",
					Category:    "creditnote",
					Flags:       []cli.Flag{creditNoteDataFlag},
					Action: func(c *cli.Context) error {
						contact, err := accounts.CreateCreditNote(creditNoteData)
						if err != nil {
							return err
						}
						fmt.Println(contact)
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
