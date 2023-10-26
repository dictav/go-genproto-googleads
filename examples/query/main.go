package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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
	accountID = flag.Int("account", 0, "ad account id (required)")
)

func main() {
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] FILE\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *accountID == 0 {
		flag.Usage()
		os.Exit(1)
	}

	f := os.Stdin

	if fn := flag.Arg(0); fn != "" {
		var err error
		if f, err = os.Open(fn); err != nil {
			log.Println(err)
			flag.Usage()
			os.Exit(1)
		}
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

	execQuery(ctx, *accountID, f, opts)
}

func execQuery(ctx context.Context, id int, r io.Reader, opts []option.ClientOption) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		log.Fatal(err)
	}

	client, err := googleads.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	req := &servicespb.SearchGoogleAdsRequest{
		CustomerId: strconv.Itoa(id),
		Query:      buf.String(),
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
