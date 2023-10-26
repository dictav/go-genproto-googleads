package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/joeshaw/envdecode"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
	ClientID       string `env:"CLIENT_ID,required"`
	ClientSecret   string `env:"CLIENT_SECRET,required"`
	RefreshToken   string `env:"REFRESH_TOKEN,required"`
	DeveloperToken string `env:"GADS_DEVELOPER_TOKEN,required"`
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

	client, err := googleads.NewCustomerClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	req := servicespb.ListAccessibleCustomersRequest{}

	res, err := client.ListAccessibleCustomers(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}

	names := strings.Join(res.GetResourceNames(), "\n")

	os.Stdout.WriteString(names)
}
