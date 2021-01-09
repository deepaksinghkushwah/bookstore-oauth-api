package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {

	assert.EqualValues(t, 24, expirationTime, "Exipration time should be 24 hours")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be nil")
	assert.EqualValues(t, "", at.AccessToken, "Access token should not have defined access token id")
	assert.True(t, at.UserID == 0, "Access token should not have an associated user id")
}

func TestIsAccessTokenExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token craeted for 3 hours from now should not be expire ")
}
