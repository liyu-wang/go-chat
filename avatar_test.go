package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/markbates/goth"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &goth.User{
		Name:      "Test User",
		AvatarURL: "",
		Provider:  "google",
		UserID:    "12345",
		Email:     "test.user@example.com",
	}
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if url != "" {
		t.Error("AuthAvatar.GetAvatarURL should return empty string when no auth data is present")
	}
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no auth data is present")
	}
	// set a value
	testUrl := "http://url-to-avatar/"
	testUser.AvatarURL = testUrl
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return an error when avatar_url is present")
	} else if url != testUrl {
		t.Errorf("AuthAvatar.GetAvatarURL returned wrong URL. Got %s, want %s", url, testUrl)
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "0bc83cb571cd1c50ba6f3e8a78ef1346"}
	url, err := gravatarAvatar.GetAvatarURL(user)
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
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
