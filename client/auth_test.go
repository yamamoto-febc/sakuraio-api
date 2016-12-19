package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSakuraAPI_Auth(t *testing.T) {

	// If APIKey isn't set , an error occurs.
	_, err := SakuraAPI.Auth()
	assert.Error(t, err)

	SakuraAPI.SetAPIKey(token, secret)
	defer SakuraAPI.ClearAPIKey()

	auth, err := SakuraAPI.Auth()
	assert.NoError(t, err)
	assert.NotNil(t, auth)

}
