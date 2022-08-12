package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromYamlFile(t *testing.T) {
	accountType, err := FromYamlFile("../test/account.yaml")
	require.NoError(t, err)
	assert.Equal(t, "account_name", accountType.Name)
	assert.Equal(t, 1, len(accountType.PropertyTypes))

	assert.Equal(t, "property_name", accountType.PropertyTypes[0].Name)
	assert.Equal(t, DataType("text"), accountType.PropertyTypes[0].DataType)
}
