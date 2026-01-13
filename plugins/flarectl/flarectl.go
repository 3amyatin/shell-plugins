package flarectl

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func FlarectlCLI() schema.Executable {
	return schema.Executable{
		Name:    "flarectl",
		Runs:    []string{"flarectl"},
		DocsURL: sdk.URL("https://github.com/cloudflare/cloudflare-go/tree/master/cmd/flarectl"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIKey,
			},
		},
	}
}
