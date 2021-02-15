package access_token

import (
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("new access token should not be expired !")

	}
	if at.AccessToken != "" {
		t.Error("new access token should have empty access token value ")

	}
	if at.UserID != 0 {
		t.Error("new access token cant have user id ")
	}
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
