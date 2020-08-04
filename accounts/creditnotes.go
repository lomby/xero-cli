package accounts

import (
	"bytes"
	"errors"

	"github.com/lomby/xero-cli/xeroclient"
)

func CreateCreditNote(creditNoteData string) (string, error) {

	if creditNoteData == "" {
		return "", errors.New("Payment data not provided")
	}

	buf := bytes.NewBuffer([]byte(creditNoteData))

	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	r, _, err := xeroclient.NewRequest("POST", "https://api.xero.com/api.xro/2.0/CreditNotes", buf, nil)

	if err != nil {
		return r, err
	}

	return r, nil

}
