package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no auth data is present")
	}
	// set a value
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]any{
		"avatar_url": testUrl,
	}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return an error when avatar_url is present")
	} else if url != testUrl {
		t.Errorf("AuthAvatar.GetAvatarURL returned wrong URL. Got %s, want %s", url, testUrl)
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]any{
		"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("GravatarAvatar.GetAvatarURL returned wrong URL. Got %s", url)
	}
}
