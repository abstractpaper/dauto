package cmd

import (
	"os"
	"os/exec"
	"testing"
)

const git_tmp_path = "/tmp/dauto_test"

func initRepo(t *testing.T) {
	cmd := exec.Command("git", "init", git_tmp_path)
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Created %s repo", git_tmp_path)
}

func deleteLocalRepo(t *testing.T) {
	err := os.RemoveAll(git_tmp_path)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Deleted %s", git_tmp_path)
}
