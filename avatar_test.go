package main

import (
	"os"
	"path/filepath"
	"testing"
)

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

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	// create avatars directory if it doesn't exist
	if _, err := os.Stat("avatars"); os.IsNotExist(err) {
		err = os.Mkdir("avatars", 0755)
		if err != nil {
			t.Fatalf("Failed to create avatars directory: %v", err)
		}
	}
	// create a dummy file to represent the avatar
	os.WriteFile(filename, []byte{}, 0777)
	defer os.Remove(filename)

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]any{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
