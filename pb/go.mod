module github.com/dictav/go-genproto-googleads/pb

go 1.16

replace google.golang.org/genproto/googleapis/ads/googleads => github.com/dictav/go-genproto-googleads/pb v0.0.0-20210701104557-0d23a41cd997

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384
	google.golang.org/genproto/googleapis/ads/googleads v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0
)
