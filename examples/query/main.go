package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	googleads "github.com/dictav/go-genproto-googleads/v8"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v8/services"
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

var account = flag.String("account", "", "ad account to execute the query. (required)")

func main() {
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()
	if *account == "" {
		flag.PrintDefaults()
		os.Exit(1)
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
		CustomerId: *account,
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
  detail_placement_view.placement,
  detail_placement_view.placement_type,
  customer.id,
  customer.descriptive_name,
  campaign.id,
  campaign.name,
  metrics.impressions,
  metrics.conversions
FROM
  detail_placement_view
WHERE
  campaign.status = ENABLED
  AND segments.date DURING LAST_7_DAYS
`
