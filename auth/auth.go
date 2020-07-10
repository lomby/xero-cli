package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

// Provider contains the oauth2 config for Xero and the context
type Provider struct {
	Config *oauth2.Config
	Ctx    context.Context
}

// NewProvider sets up a new Xero Provider containing the oauth2 config and the context
func NewProvider() *Provider {

	endpoint := oauth2.Endpoint{
		AuthURL:  "https://login.xero.com/identity/connect/authorize",
		TokenURL: "https://identity.xero.com/connect/token",
	}

	return &Provider{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("XERO_KEY"),
			ClientSecret: os.Getenv("XERO_SECRET"),
			Endpoint:     endpoint,
			RedirectURL:  "http://localhost:8080/callback",
			Scopes:       strings.Split(os.Getenv("XERO_SCOPES"), ","),
		},
		Ctx: context.Background(),
	}

}

// HandleCallback for initial manual auth. The auth code will be redirected here where we'll request the oauth2 token and save it locally.
func (p *Provider) HandleCallback(res http.ResponseWriter, req *http.Request) {

	// Fetch the code returned from Xero auth
	code := req.FormValue("code")
	token, err := p.Config.Exchange(p.Ctx, code)

	if err != nil {
		fmt.Println(err)
	}

	// Save our initial token to credentials.json
	storeToken(*token)
}

// GetAuthURL fetches the xero url for manual auth
func (p *Provider) GetAuthURL() {

	// Fetch the Xero Auth URL to manually auth with user login in browser
	authURL := p.Config.AuthCodeURL("UNIQUE_STATE_STRING")
	fmt.Println(authURL)
}

// GetToken fetched oauth2 token from our local file
func (p *Provider) GetToken() oauth2.Token {

	tokenFile, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	var token oauth2.Token
	json.Unmarshal(tokenFile, &token)

	// Check the token hasn't expired. If it has, refresh it.
	if token.Expiry.Before(time.Now()) {
		token = p.refreshToken(token)
	}

	return token

}

// Refresh for the AuthToken
func (p *Provider) refreshToken(token oauth2.Token) oauth2.Token {

	tokenSource := p.Config.TokenSource(p.Ctx, &token)
	newToken, err := tokenSource.Token()

	if err != nil {
		fmt.Println(err)
	}

	if newToken.AccessToken != token.AccessToken {
		token = *newToken
		storeToken(token)
	}

	return token
}

// StoreToken is a helper function that stores the oauth2 token to a loclal file
func storeToken(token oauth2.Token) {
	file, err := json.MarshalIndent(token, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("credentials.json", file, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
