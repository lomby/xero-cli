package xeroclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/lomby/xero-cli/auth"
)

// NewRequest is a wrapper around http.NewRequest where we add in Xero Provider & Token and Tenant header
// Accepts same arguments as http.NewRequest
//
// Returns repsonse body as string
func NewRequest(method string, url string, body io.Reader) (response string, statusCode int, err error) {

	provider := auth.NewProvider()
	token := provider.GetToken()

	client := provider.Config.Client(provider.Ctx, &token)

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("xero-tenant-id", os.Getenv("UK_TENANT_ID"))

	resp, err := client.Do(request)

	if err != nil {
		return "", resp.StatusCode, err
	}

	defer resp.Body.Close()

	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

	return bodyString, resp.StatusCode, nil

}

func GetTenants(res http.ResponseWriter, req *http.Request) {

	r, code, err := NewRequest("GET", "https://api.xero.com/connections", nil)

	if err != nil {
		fmt.Println(code, err)
	}

	fmt.Println(code, r)

}
