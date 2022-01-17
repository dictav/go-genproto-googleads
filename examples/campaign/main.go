package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	servicespb "github.com/dictav/go-genproto-googleads/pb/v9/services"
	googleads "github.com/dictav/go-genproto-googleads/v9"
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

var (
	accountID  = flag.Int("account", 0, "ad account id (required)")
	campaignID = flag.Int("campaign", 0, "campaign id")
)

func main() {
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()
	if *accountID == 0 {
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

	if *campaignID > 0 {
		show(ctx, *campaignID, opts)
		return
	}

	list(ctx, *accountID, opts)
}

func show(ctx context.Context, id int, opts []option.ClientOption) {
	client, err := googleads.NewCampaignClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	req := &servicespb.GetCampaignRequest{
		ResourceName: fmt.Sprintf("customers/%d/campaigns/%d", *accountID, *campaignID),
	}

	res, err := client.GetCampaign(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err := enc.Encode(res); err != nil {
		log.Fatalln(err)
	}
}

func list(ctx context.Context, id int, opts []option.ClientOption) {
	client, err := googleads.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	req := &servicespb.SearchGoogleAdsRequest{
		CustomerId: strconv.Itoa(id),
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
  campaign.id,
	campaign.name
FROM
  campaign
`
