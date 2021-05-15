module github.com/dictav/go-genproto-googleads/examples

go 1.16

replace (
	github.com/dictav/go-genproto-googleads => ../
	google.golang.org/genproto/googleapis/ads/googleads => ../pb
)

require (
	github.com/dictav/go-genproto-googleads v0.0.0-00010101000000-000000000000
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	google.golang.org/api v0.46.0
	google.golang.org/grpc v1.37.1
)
