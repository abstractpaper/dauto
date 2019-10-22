package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

// Test FileExists function for a file that doesn't exist
func TestFileExistsNotFound(t *testing.T) {
	path := "/tmp/dauto_unknown_file"
	// File should not exist
	exists, _ := FileExists(path)
	if exists {
		t.Fatalf("Expected: file does not exist; Got: file exists at %s", path)
	}
}

// Test FileExists function for a file that exists
func TestFileExistsFound(t *testing.T) {
	path := "/tmp/dauto_unknown_file"

	// Write file at path
	err := ioutil.WriteFile(path, []byte("test"), 0744)
	if err != nil {
		t.Fatal(err)
	}
	// Make sure to delete file at the end
	defer func() {
		err = os.Remove(path)
		if err != nil {
			t.Fatal(err)
		}
	}()

	// File should exist
	exists, _ := FileExists(path)
	if !exists {
		t.Fatalf("Expected: file exists; Got: file does not exist at %s", path)
	}
}
