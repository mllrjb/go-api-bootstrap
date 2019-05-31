## Getting Started

`go get github.com/logrhythm/go-api-bootstrap` to clone the bootstrap and get started.

This uses go modules, which works fine with go@1.12+ and go-swagger@0.19.0+. Go modules is automatically enabled if your project source is *outside* `GOPATH`, otherwise you'll need to `export GO111MODULE=on` for it to work.

## Installing Dependencies

This project uses go modules, which is available in go@1.11+.

```
go mod download
```

Add new dependencies using `go get -u {package}`.

## Generating go-swagger server

https://github.com/go-swagger/go-swagger/issues/1724#issuecomment-469335593

To generate the server, run:

```
go generate
```

Thanks to folks at Netlify for their approach to [go-swagger + modules](
https://github.com/go-swagger/go-swagger/issues/1724#issuecomment-469335593)!

## Build the Server

```
go build -o out/API
```

## Run the Server

```
HTTP_PORT=80 BIND_ADDRESS=0.0.0.0 ./out/API
```
