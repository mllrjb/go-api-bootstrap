# Getting Started

## Prerequisites

* go@1.12+
* [gobin](https://github.com/myitcv/gobin)

## Installing Dependencies

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

(ignore messages about adding additional packages in your GOPATH)

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
