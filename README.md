# go-genproto-googleads

> **IMPORTANT** This is experimental and there is no guarantee for use this library. However, if you are interested in trying out the Google Ads API from Go, this repository will help you.
>
> If you hope the official Google Ads API for Go, please comment in the following:
>
> https://groups.google.com/g/adwords-api/c/pZJyu8Ih3GI/m/1xZlFz4bAwAJ

This is Google Ads API Go Client. The sources is auto generated from the proto files in [googleapis/googleapis].


## Install

```sh
$ go mod edit -require github.com/dictav/go-genproto-googleads@v0.20220116.1
```

## How to generate

I forked [googleapis/googleapis] and wrote patches for building Go Client of Google Ads API [here](https://github.com/dictav/googleapis).

```sh
$ make all
```

Enjoy!

## Alternates

- https://github.com/googleapis/google-cloud-go
- https://github.com/googleapis/go-genproto


[googleapis/googleapis]: https://github.com/googleapis/googleapis/
