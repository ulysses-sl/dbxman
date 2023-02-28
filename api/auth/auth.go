package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	dbxman "github.com/ulysses-sl/dbxman/api"
	"golang.org/x/oauth2"
	"os"
)

type pkceCode string

const pkceChallangeMethod = "S256"
const dbxAppKeyEnv = "DROPBOX_APP_KEY"
const dbxAuthURI = "/oauth2/authorize"
const dbxTokenURI = "/oauth2/token"

type DbxAuth interface {
	Authorize() error
}

type dbxOAuth struct {
	endpoint oauth2.Endpoint
}

func Generate() {
	rand.Reader
	sha256.New()
}

func NewDbxOAuth(baseURL *dbxman.BaseURL) DbxAuth {
	authURL := baseURL.OAuth.JoinPath(dbxAuthURI)
	tokenURL := baseURL.Api.JoinPath(dbxTokenURI)
	auth := dbxOAuth{
		endpoint: oauth2.Endpoint{
			AuthURL:  authURL.String(),
			TokenURL: tokenURL.String(),
		},
	}
	return &auth
}

func (auth *dbxOAuth) Authorize() error {
	appKey := os.Getenv(dbxAppKeyEnv)
	if len(appKey) == 0 {
		return fmt.Errorf("the app key for Dropbox API should be "+
			"supplied in the environment variable %s", dbxAppKeyEnv)
	}
	ctx := context.Background()
	conf := oauth2.Config{
		ClientID: appKey,
	}
	return nil
}
