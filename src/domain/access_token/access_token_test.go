package access_token

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(),"brand new access token sholud not be nil" )
	assert.EqualValues(t,"", at.AccessToken, "new access token sholud not have defined access token id" )
	assert.True(t,at.UserId == 0, "new acces token should not have an associated user id" )

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(),"default access token should be expired by default" )
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(),"access token expiring three hours from new should not be expired" )
}