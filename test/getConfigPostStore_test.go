package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/config"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/poststore"
	"github.com/stretchr/testify/assert"
)

func TestGetConfiguration(t *testing.T) {
	ps, err := poststore.New()
	assert.Nil(t, err)
	assert.NotNil(t, ps)

	testConfig := &config.Config{
		ID:      "test-id",
		Version: "1",
		Name:    "Test Configuration",
	}

	err = ps.AddConfiguration(context.Background(), testConfig)
	assert.Nil(t, err)

	retrievedConfig, err := ps.GetConfiguration(context.Background(), testConfig.ID, testConfig.Version)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedConfig)
	assert.Equal(t, testConfig.ID, retrievedConfig.ID)
	assert.Equal(t, testConfig.Version, retrievedConfig.Version)
	assert.Equal(t, testConfig.Name, retrievedConfig.Name)

	fmt.Println("TestGetConfiguration - Test Started")
	fmt.Println("TestGetConfiguration - Added Configuration:", testConfig)
	fmt.Println("TestGetConfiguration - Retrieved Configuration:", retrievedConfig)
	fmt.Println("TestGetConfiguration - ID Matched:", testConfig.ID == retrievedConfig.ID)
	fmt.Println("TestGetConfiguration - Version Matched:", testConfig.Version == retrievedConfig.Version)
	fmt.Println("TestGetConfiguration - Name Matched:", testConfig.Name == retrievedConfig.Name)
	fmt.Println("TestGetConfiguration - Test Finished")
}
