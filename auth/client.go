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

type Provider struct {
	Config *oauth2.Config
	Ctx    context.Context
}

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

func (p *Provider) GetAuthURL() {

	// Fetch the Xero Auth URL to manually auth with user login in browser
	authURL := p.Config.AuthCodeURL("UNIQUE_STATE_STRING")
	fmt.Println(authURL)
}

func (p *Provider) GetToken() oauth2.Token {

	tokenFile, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		fmt.Println(err)
	}

	var token oauth2.Token
	json.Unmarshal(tokenFile, &token)

	if token.Expiry.Before(time.Now()) {
		token = p.refreshToken(token)
	}

	return token

}

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
