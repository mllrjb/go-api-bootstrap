## Getting Started

`go get github.com/logrhythm/go-api-bootstrap` to clone the bootstrap and get started.

This uses go modules, which works fine with go@1.12+ and go-swagger@0.19.0+. Go modules is automatically enabled if your project source is *outside* `GOPATH`, otherwise you'll need to `export GO111MODULE=on` for it to work.

## Installing Dependencies

This project uses go modules, which is available in go@1.11+.

Install dependencies with `go get`.

Add a dependency using `go get -u {package}`.

## Generating go-swagger server

To generate the server, first install the `go-swagger` tools:

```
brew tap go-swagger/go-swagger
brew install go-swagger
```

```
swagger generate server \
  -A api \
  -P auth.UserPrincipal \
  -t generated \
  --exclude-main \
  -s swagger \
  -f swagger.yaml
```

## Build the Server

Enable go modules and build the API.

```
go build -o out/API
```

## Run the Server

```
HTTP_PORT=80 BIND_ADDRESS=0.0.0.0 ./API
```
