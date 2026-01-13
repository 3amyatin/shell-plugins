package flarectl

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestAPIKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Email:  "user@example.com",
				fieldname.APIKey: "test-api-key-37-chars-long-here123",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"CF_API_EMAIL": "user@example.com",
					"CF_API_KEY":   "test-api-key-37-chars-long-here123",
				},
			},
		},
	})
}

func TestAPIKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"CF_API_EMAIL": "user@example.com",
				"CF_API_KEY":   "test-api-key-37-chars-long-here123",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Email:  "user@example.com",
						fieldname.APIKey: "test-api-key-37-chars-long-here123",
					},
				},
			},
		},
		"config file": {
			Files: map[string]string{},
			ExpectedCandidates: []sdk.ImportCandidate{},
		},
	})
}
