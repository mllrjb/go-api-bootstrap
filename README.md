## Getting Started

`go get github.com/logrhythm/go-api-bootstrap` to clone the bootstrap and get started.

Note that you'll want to move it to another repository and location on disk in order to customize and publich it. Even though this project is using go modules, it isn't fully supported in the `go-swagger` toolchain, so your project will still need to be located under your `GOPATH`.

## Installing Dependencies

This project uses go modules, which is available in go@1.11+.

Install dependencies with `GO111MODULE=on go get`.

Add a dependency using `GO111MODULE=on go get -u {package}`.

## Generating go-swagger server

To generate the server, first install the `go-swagger` tools:

```
brew tap go-swagger/go-swagger
brew install go-swagger
```

When generating the server, you'll need to ensure you disable go modules, as `go-swagger` relies on a version of `go import` that does not support modules yet. See [this issue](https://github.com/go-swagger/go-swagger/issues/1882) for details.

```
GO111MODULE=off swagger generate server \
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
GO111MODULE=on go build -o out/API
```

## Run the Server

```
HTTP_PORT=80 BIND_ADDRESS=0.0.0.0 ./API
```
