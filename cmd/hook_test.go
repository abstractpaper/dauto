package cmd

import (
	"testing"
)

func TestHook(t *testing.T) {
	initRepo(t)
	hook(git_tmp_path)
	deleteLocalRepo(t)
}
