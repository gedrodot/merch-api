package jwt

import "testing"

func Test(t *testing.T) {
	testUserId := 52
	token, err := GenerateToken(testUserId)
	if err != nil {
		t.Fatalf("token generation error %v", err)
	}
	t.Logf("token: %s", token)
	parsedId, err := ParseToken(token)
	if err != nil {
		t.Fatalf("error parse token %v", err)
	}
	if parsedId != testUserId {
		t.Errorf("kaka")
	}
	t.Log(parsedId)
	t.Log(secretKey)
}
