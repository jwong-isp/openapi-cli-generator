package main

import (
	"github.com/danielgtaylor/openapi-cli-generator/cli"
	"github.com/danielgtaylor/openapi-cli-generator/credentials"
)

//go:generate openapi-cli-generator generate openapi.yaml

func main() {
	cli.Init(&cli.Config{
		AppName:   "example",
		EnvPrefix: "EXAMPLE",
		Version:   "1.0.0",
	})
	credentials.RegisterAuth(&credentials.PKCE{
		clientID: "cid"
		issuerEndpoint: "https://example.com"
	})

	openapiRegister(false)

	cli.Root.Execute()
}
