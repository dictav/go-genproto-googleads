package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	servicespb "github.com/dictav/go-genproto-googleads/pb/v14/services"
	googleads "github.com/dictav/go-genproto-googleads/v14"
)

const (
	gadsAddr = "googleads.googleapis.com:443"
	scope    = "https://www.googleapis.com/auth/adwords"
)

var env struct {
	ClientID        string `env:"CLIENT_ID,required"`
	ClientSecret    string `env:"CLIENT_SECRET,required"`
	RefreshToken    string `env:"REFRESH_TOKEN,required"`
	DeveloperToken  string `env:"GADS_DEVELOPER_TOKEN,required"`
	LoginCustomerID string `env:"LOGIN_CUSTOMER_ID,required"`
}

func main() {
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	t := &oauth2.Token{RefreshToken: env.RefreshToken, TokenType: "Bearer"}
	conf := &oauth2.Config{
		ClientID:     env.ClientID,
		ClientSecret: env.ClientSecret,
		Scopes:       []string{scope},
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080",
	}

	opts := []option.ClientOption{
		option.WithTokenSource(conf.TokenSource(ctx, t)),
		option.WithEndpoint(gadsAddr),
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "developer-token", env.DeveloperToken)
	ctx = metadata.AppendToOutgoingContext(ctx, "login-customer-id", env.LoginCustomerID)

	client, err := googleads.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	req := &servicespb.SearchGoogleAdsRequest{
		CustomerId: env.LoginCustomerID,
		Query:      query,
	}

	it := client.Search(ctx, req)
	enc := json.NewEncoder(os.Stdout)

	for {
		row, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		if err := enc.Encode(row); err != nil {
			log.Fatalln(err)
		}
	}
}

const query = `
SELECT
	customer_client.id,
	customer_client.descriptive_name,
	customer_client.manager
FROM
	customer_client
WHERE
	customer_client.status = ENABLED
`
