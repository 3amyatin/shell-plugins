package flarectl

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIKey() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIKey,
		DocsURL:       sdk.URL("https://developers.cloudflare.com/fundamentals/api/get-started/keys/"),
		ManagementURL: sdk.URL("https://dash.cloudflare.com/profile/api-tokens"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Email,
				MarkdownDescription: "Email address associated with your Cloudflare account.",
				Optional:            false,
			},
			{
				Name:                fieldname.APIKey,
				MarkdownDescription: "Global API Key for your Cloudflare account. Found in your [Cloudflare profile](https://dash.cloudflare.com/profile/api-tokens).",
				Secret:              true,
				Optional:            false,
				Composition: &schema.ValueComposition{
					Length: 37,
					Charset: schema.Charset{
						Lowercase: true,
						Uppercase: true,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(map[string]sdk.FieldName{
			"CF_API_EMAIL": fieldname.Email,
			"CF_API_KEY":   fieldname.APIKey,
		}),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(map[string]sdk.FieldName{
				"CF_API_EMAIL": fieldname.Email,
				"CF_API_KEY":   fieldname.APIKey,
			}),
		),
	}
}
