package cmd

import (
	"fmt"
	"testing"

	"github.com/abstractpaper/dauto/utils"
)

// Test hooking a repo
func TestHook(t *testing.T) {
	initRepo(t)
	defer deleteLocalRepo(t)

	// hook repo
	err := hook(git_tmp_path)
	if err != nil {
		t.Fatal(err)
	}

	// check if hook was created
	path := fmt.Sprintf("%s/.git/hooks/pre-commit", git_tmp_path)
	exists, err := utils.FileExists(path)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("pre-commit hook was not created at: %s", path)
	}
}

// Test hooking a repo that's already hooked
func TestHookExists(t *testing.T) {
	initRepo(t)
	defer deleteLocalRepo(t)

	// hook repo
	err := hook(git_tmp_path)
	if err != nil {
		t.Fatal(err)
	}

	// check if initial hook was created
	path := fmt.Sprintf("%s/.git/hooks/pre-commit", git_tmp_path)
	exists, err := utils.FileExists(path)
	if err != nil {
		t.Fatal(err)
	}
	// hook function was not successful in creating a hook
	if !exists {
		t.Fatalf("pre-commit hook was not created at: %s", path)
	}

	// try to hook repo again
	err = hook(git_tmp_path)
	if err == nil {
		t.Fatal("Hook overwrote an existing hook successfully; error expected.")
	}
}
