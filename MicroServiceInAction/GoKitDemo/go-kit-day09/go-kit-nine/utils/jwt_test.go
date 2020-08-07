package utils

import "testing"

func TestGenerParseToken(t *testing.T) {

	token, err := GenerToken("yuzihan", 1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("token :", token)

	result, err := ParseToken(token)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("result :", result)

}
