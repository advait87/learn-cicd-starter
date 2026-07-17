package auth

import (
	"net/http"
	"testing"
)

func TestMalinformedHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{"breadf": []string{"ApiKey"}})
	if err == nil {
		t.Error("Expected error when passing a malinformed header")
	}
}

func TestNormalHeader(t *testing.T) {
	apiKey, err := GetAPIKey(http.Header{"Authorization": []string{"ApiKey slkdfj"}})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if apiKey != "slkdfj" {
		t.Errorf("Expected api key to be slkdfj, got %s", apiKey)
	}
}
