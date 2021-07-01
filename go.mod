module github.com/dictav/go-genproto-googleads

go 1.16

replace google.golang.org/genproto/googleapis/ads/googleads => ./pb

require (
	cloud.google.com/go v0.85.0
	github.com/googleapis/gax-go/v2 v2.0.5
	google.golang.org/api v0.50.0
	google.golang.org/genproto v0.0.0-20210630183607-d20f26d13c79
	google.golang.org/genproto/googleapis/ads/googleads v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
