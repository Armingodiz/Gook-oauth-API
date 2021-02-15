package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "new access token should not be expired !")
	assert.EqualValues(t, "", at.AccessToken, "new access token should have empty access token value ")
	assert.True(t, at.UserID == 0, "new access token cant have user id ")
}
func TestAccessToken_IsExpired(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("new access token must not be expired")

	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token created 3 hours from now should not be expired !")
	}
}
