package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_NoHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatal("expected error when no Authorization header is present, got nil")
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer sometoken")
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error for malformed Authorization header, got nil")
	}
}

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	want := "my-secret-key"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
