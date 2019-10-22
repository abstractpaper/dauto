package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestHook(t *testing.T) {
	initRepo(t)
	hook(git_tmp_path)

	path := fmt.Sprintf("%s/.git/hooks/pre-commit", git_tmp_path)
	info, err := os.Stat(path)
	if os.IsNotExist(err) || info.IsDir() {
		t.Fatalf("pre-commit hook was not created at: %s", path)
	}

	deleteLocalRepo(t)
}
