module github.com/dictav/go-genproto-googleads

go 1.16

replace google.golang.org/genproto/googleapis/ads/googleads => ./pb

require (
	cloud.google.com/go v0.81.0
	github.com/googleapis/gax-go/v2 v2.0.5
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384
	google.golang.org/genproto/googleapis/ads/googleads v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0
)
