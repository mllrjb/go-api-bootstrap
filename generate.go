package main

//go:generate gobin -m -run github.com/go-swagger/go-swagger/cmd/swagger generate server -A api -P auth.UserPrincipal -t generated --exclude-main -s swagger -f swagger.yaml
