package test

import (
	"context"
	"fmt"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/config"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/poststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddConfiguration(t *testing.T) {
	ps, err := poststore.New()
	assert.Nil(t, err)
	assert.NotNil(t, ps)

	testConfig := &config.Config{
		ID:      "test-id",
		Version: "1",
		Name:    "Test Configuration",
	}
	fmt.Println("Adding configuration:", testConfig)

	err = ps.AddConfiguration(context.Background(), testConfig)
	assert.Nil(t, err)

	fmt.Println("Retrieving configuration with ID:", testConfig.ID, "and version:", testConfig.Version)

	retrievedConfig, err := ps.GetConfiguration(context.Background(), testConfig.ID, testConfig.Version)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedConfig)
	assert.Equal(t, testConfig.ID, retrievedConfig.ID)
	assert.Equal(t, testConfig.Version, retrievedConfig.Version)
	assert.Equal(t, testConfig.Name, retrievedConfig.Name)
	fmt.Println("Retrieved configuration:", retrievedConfig)

}
