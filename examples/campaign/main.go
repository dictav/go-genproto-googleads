package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"strconv"

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
		show(ctx, *accountID, *campaignID, opts)
		return
	}

	list(ctx, *accountID, opts)
}

func show(ctx context.Context, accountID, campaignID int, opts []option.ClientOption) {
	log.Println("WARNING: it does not include BiddingStrategy")

	client, err := googleads.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	_ = campaignID

	req := &servicespb.SearchGoogleAdsRequest{
		CustomerId: strconv.Itoa(accountID),
		Query:      queryDetail + strconv.Itoa(campaignID),
	}

	it := client.Search(ctx, req)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	row, err := it.Next()
	if err != nil {
		log.Fatalln(err)
	}

	c := row.GetCampaign()
	log.Printf("bidding_strategy: %#v\n", c.GetBiddingStrategy())
	log.Printf("bidding_strategy_type: %#v\n", c.GetBiddingStrategyType())

	if err := enc.Encode(c); err != nil {
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
	campaign.name,
	campaign.advertising_channel_type,
	campaign.advertising_channel_sub_type,
	campaign.status,
	campaign.serving_status,
	campaign.start_date,
	campaign.end_date,
	ad_group.id,
	ad_group.name,
	ad_group.type
FROM
	ad_group
WHERE
	campaign.status IN (ENABLED, PAUSED)
AND
	ad_group.status = ENABLED
ORDER BY
	campaign.end_date DESC
`

const queryDetail = `
SELECT
	campaign.id,
	campaign.name,
	campaign.status,
	campaign.serving_status,
	campaign.ad_serving_optimization_status,
	campaign.advertising_channel_type,
	campaign.advertising_channel_sub_type,
	campaign.experiment_type,
	campaign.base_campaign,
	campaign.campaign_budget,
	campaign.bidding_strategy_type,
	campaign.start_date,
	campaign.end_date,
	campaign.payment_mode,
	campaign.network_settings.target_google_search,
	campaign.network_settings.target_search_network,
	campaign.network_settings.target_content_network,
	campaign.network_settings.target_partner_search_network,
	campaign.geo_target_type_setting.positive_geo_target_type,
	campaign.geo_target_type_setting.negative_geo_target_type,
	campaign.app_campaign_setting.app_id,
	campaign.app_campaign_setting.app_store,
	campaign.app_campaign_setting.bidding_strategy_goal_type,
	campaign.bidding_strategy,
	campaign.bidding_strategy_type
FROM
	campaign
WHERE
	campaign.id =`
