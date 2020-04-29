package tfe

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModulesCreate(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	testOrg, testOrgCleanup := createOrganization(t, client)
	defer testOrgCleanup()

	optionsModule := ModuleCreateOptions{
		Name: *String(randomString(t)),
		Provider: "random",
	}

	t.Run("creating a module", func(t *testing.T) {
		m, err := client.Registry.CreateModule(ctx, testOrg.Name, optionsModule)
		require.NoError(t, err)
		assert.Equal(t, optionsModule.Name, m.Name)
		assert.Equal(t, optionsModule.Provider, m.Provider)
	})

	t.Run("creating a module version", func(t *testing.T) {
		optionsModuleVersion := ModuleCreateVersionOptions{
			Version: "1.2.3",
		}

		mv, err := client.Registry.CreateModuleVersion(ctx, testOrg.Name, optionsModule.Name, optionsModule.Provider, optionsModuleVersion)
		require.NoError(t, err)
		assert.Equal(t, optionsModuleVersion.Version, mv.Version)
	})
}
