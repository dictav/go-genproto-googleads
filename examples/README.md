# Examples

## Prepare environment variables

```
$ export CLIENT_ID=*****
$ export CLIENT_SECRET=*****
$ export GADS_DEVELOPER_TOKEN=*****
$ export LOGIN_CUSTOMER_ID=0000000000
$ export REFRESH_TOKEN=*****
```

## Usage

```
$ go run ./customer/main.go
```

```
$ go run ./account/main.go
```

```
$ go run ./campaign/main.go -account=<AD_ACCOUNT_ID> -campaign=<CAMPAIGN_ID>
```

```
$ go run ./query/main.go -account=<AD_ACCOUNT_ID> FILE
```
