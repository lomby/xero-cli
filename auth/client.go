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

type credentials struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Provider struct {

Config *oauth2.Config
Ctx context.Context

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


	code := req.FormValue("code")
	token, err := p.Config.Exchange(p.Ctx, code)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)

	file, _ := json.MarshalIndent(token, "", " ")

	_ = ioutil.WriteFile("credentials.json", file, 0644)
}

func (p *Provider) GetAuthURL() {

	authURL := p.Config.AuthCodeURL("UNIQUE_STATE_STRING")
	fmt.Println(authURL)
}