package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/lomby/xero-cli/xeroclient"
)

func GetContact(contactID string, search string) (r string, err error) {

	if contactID == "" && search == "" {
		return "", errors.New("Neither a Contact ID or search not provided")
	}

	var code int

	if contactID != "" {
		r, code, err = xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Contacts/"+contactID, nil, nil)
	}

	if search != "" {

		var searchString strings.Builder

		var searchMap map[string]string
		err := json.Unmarshal([]byte(search), &searchMap)

		if err != nil {
			fmt.Println(errors.New("Invalid search data provided"))
		}

		for key, value := range searchMap {
			searchString.WriteString(key + "==\"" + value + "\"")
		}

		r, code, err = xeroclient.NewRequest("GET", "https://api.xero.com/api.xro/2.0/Contacts?where="+url.QueryEscape(searchString.String()), nil, nil)
	}

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil

}

func CreateContact(contactData string) (string, error) {

	if contactData == "" {
		return "", errors.New("Contact data not provided")
	}

	r, code, err := xeroclient.NewRequest("POST", "https://api.xero.com/api.xro/2.0/Contacts", bytes.NewBuffer([]byte(contactData)), nil)

	if err != nil || code != 200 {
		return "", err
	}

	return r, nil
}
