package accounts

import (
	"fmt"
	"net/http"

	"github.com/lomby/xero-cli/xeroclient"
)

func GetInvoice(res http.ResponseWriter, req *http.Request) {

	invoiceID := req.FormValue("id")

	if invoiceID == "" {
		fmt.Println("Invoice ID not provided")
		return
	}

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/"+invoiceID, nil)

	if err != nil || code != 200 {
		fmt.Println(code, err, r)
		return
	}

	fmt.Println(code, r)
	return
}
