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
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.APIKey: "kXbfANDLo2LbLvE8TZnvc3HYbF",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"FLARECTL_API_KEY": "kXbfANDLo2LbLvE8TZnvc3HYbF",
				},
			},
		},
	})
}

func TestAPIKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"FLARECTL_API_KEY": "kXbfANDLo2LbLvE8TZnvc3HYbF",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.APIKey: "kXbfANDLo2LbLvE8TZnvc3HYbF",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in flarectl/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "kXbfANDLo2LbLvE8TZnvc3HYbF",
			// 		},
			// 	},
			},
		},
	})
}
