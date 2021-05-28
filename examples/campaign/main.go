package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"

	googleads "github.com/dictav/go-genproto-googleads/v7"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v7/services"
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
	campaignID = flag.Int("campaign", 0, "campaign id (required)")
)

func main() {
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()
	if *accountID == 0 || *campaignID == 0 {
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
