package accounts

import (
	"bytes"
	"errors"

	"github.com/lomby/xero-cli/xeroclient"
)

// GetInvoice fetches a single invoice when provided with a Xero InvoiceID
func GetInvoice(invoiceID string, pdf bool) (string, error) {

	if invoiceID == "" {
		return "", errors.New("Invoice ID not provided")
	}

	var headers = make(map[string]string)

	if pdf {
		headers["Accept"] = "application/pdf"
	}

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/"+invoiceID, nil, headers)

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil
}

// GetInvoices fetches a all invoices when provided with a CustomerID
func GetInvoices(contactID string) (string, error) {

	if contactID == "" {
		return "", errors.New("Contact ID not provided")
	}

	r, _, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices?ContactIDs="+contactID, nil, nil)

	if err != nil {
		return "", err
	}

	return r, nil
}

// GetInvoices fetches a all invoices when provided with a CustomerID
func CreateInvoice(invoiceData string) (string, error) {

	if invoiceData == "" {
		return "", errors.New("Invoice data not provided")
	}

	buf := bytes.NewBuffer([]byte(invoiceData))

	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	r, _, err := xeroclient.NewRequest("POST", "https://api.xero.com/api.xro/2.0/Invoices", buf, nil)

	if err != nil {
		return r, err
	}

	return r, nil
}

// GetInvoiceLink fetches the online url for an invoice using the invoice id
func GetInvoiceLink(invoiceID string) (string, error) {

	if invoiceID == "" {
		return "", errors.New("Invoice ID not provided")
	}

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/"+invoiceID+"/OnlineInvoice", nil, nil)

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil
}
