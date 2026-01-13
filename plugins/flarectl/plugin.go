package flarectl

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "flarectl",
		Platform: schema.PlatformInfo{
			Name:     "Cloudflare",
			Homepage: sdk.URL("https://cloudflare.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			CloudflareCLI(),
		},
	}
}
