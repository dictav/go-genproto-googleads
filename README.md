# go-genproto-googleads

> **IMPORTANT** This is experimental and there is no guarantee for use this library. However, if you are interested in trying out the Google Ads API from Go, this repository will help you.
>
> If you hope the official Google Ads API for Go, please comment in the following:
>
> https://groups.google.com/g/adwords-api/c/pZJyu8Ih3GI/m/1xZlFz4bAwAJ

```sh
go get github.com/dictav/go-genproto-googleads
```

This is Google Ads API Go Client. The sources is auto generated from the proto files in [googleapis/googleapis].


## How to generate

I forked [googleapis/googleapis] and wrote patches for building Go Client of Google Ads API [here](https://github.com/dictav/googleapis).

```sh
bazel build //google/ads/googleads/v7:gapi-ads-googleads-v7-go
```

Enjoy!

## Alternates

- https://github.com/googleapis/google-cloud-go
- https://github.com/googleapis/go-genproto


[googleapis/googleapis]: https://github.com/googleapis/googleapis/
