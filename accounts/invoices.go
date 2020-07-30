package accounts

import (
	"errors"

	"github.com/lomby/xero-cli/xeroclient"
)

// GetInvoice fetches a single invoice when provided with a Xero InvoiceID
func GetInvoice(invoiceID string) (string, error) {

	if invoiceID == "" {
		return "", errors.New("Invoice ID not provided")
	}

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/"+invoiceID, nil)

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

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices?ContactIDs="+contactID, nil)

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil
}

// GetInvoiceLink fetches the online url for an invoice using the invoice id
func GetInvoiceLink(invoiceID string) (string, error) {

	if invoiceID == "" {
		return "", errors.New("Invoice ID not provided")
	}

	r, code, err := xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Invoices/"+invoiceID+"/OnlineInvoice", nil)

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil
}
