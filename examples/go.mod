module github.com/dictav/go-genproto-googleads/examples/v8

go 1.16

replace (
	github.com/dictav/go-genproto-googleads => ../
	google.golang.org/genproto/googleapis/ads/googleads => ../pb
)

require (
	github.com/dictav/go-genproto-googleads v0.0.0-00010101000000-000000000000
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914
	google.golang.org/api v0.50.0
	google.golang.org/genproto/googleapis/ads/googleads v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
)
