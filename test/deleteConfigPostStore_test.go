package test

import (
	"context"
	"fmt"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/config"
	"github.com/anna02272/AlatiZaRazvojSoftvera2023-projekat/poststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteConfiguration(t *testing.T) {
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

	err = ps.DeleteConfiguration(context.Background(), testConfig.ID, testConfig.Version)
	assert.Nil(t, err)

	_, err = ps.GetConfiguration(context.Background(), testConfig.ID, testConfig.Version)
	assert.NotNil(t, err)

	fmt.Println("TestDeleteConfiguration - Test Started")
	fmt.Println("TestDeleteConfiguration - Configuration Deleted:", testConfig)
	fmt.Println("TestDeleteConfiguration - Configuration Still Exists:", err != nil)
	fmt.Println("TestDeleteConfiguration - Test Finished")
}
