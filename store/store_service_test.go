package store

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	os.Setenv("APP_ENV", "testing")
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialUrl := "https://example.com"
	userID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX24"

	SaveUrlMapping(shortURL, initialUrl, userID)

	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialUrl, retrievedUrl)
}
