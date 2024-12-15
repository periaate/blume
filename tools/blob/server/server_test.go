package server

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/periaate/blob"
)

func setupTestServer() *httptest.Server {
	return httptest.NewServer(NewServer())
}

func TestClientServerIntegration(t *testing.T) {
	blob.SetIndex("./testing")
	// Start test server
	testServer := setupTestServer()
	defer testServer.Close()

	client := NewClient(testServer.URL)

	// Test Set operation
	data := strings.NewReader("integration test content")
	err := client.Set("test-bucket", "test-blob", data, blob.PLAIN)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	// Test Get operation
	reader, contentType, err := client.Get("test-bucket", "test-blob")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	defer reader.Close()

	content, _ := io.ReadAll(reader)
	if string(content) != "integration test content" {
		t.Errorf("expected content 'integration test content', got '%s'", string(content))
	}
	if contentType != blob.PLAIN {
		t.Errorf("expected content type 'text/plain', got '%s'", contentType)
	}

	// Test Delete operation
	err = client.Delete("test-bucket", "test-blob")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Verify deletion
	_, _, err = client.Get("test-bucket", "test-blob")
	if err == nil {
		t.Fatalf("expected Get to fail after deletion, but it succeeded")
	}
}
