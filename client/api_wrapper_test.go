package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSakuraAPI(t *testing.T) {
	client := SakuraAPI.(*SakuraAPIClient)
	assert.NotNil(t, client)
	assert.Nil(t, client.basicAuthInfoWriter)
}

func TestSakuraAPI_SetAPIKey(t *testing.T) {
	client := SakuraAPI.(*SakuraAPIClient)
	client.SetAPIKey("token", "secret")
	assert.NotNil(t, client.basicAuthInfoWriter)
}

func TestSakuraAPI_ClearAPIKey(t *testing.T) {
	client := SakuraAPI.(*SakuraAPIClient)
	client.ClearAPIKey()
	assert.Nil(t, client.basicAuthInfoWriter)
}
