package cmd

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRun(t *testing.T) {
	initRepo(t)
	hook(git_tmp_path)

	// setup dauto.json
	json := `["echo rgr"]`
	json_path := fmt.Sprintf("%s/dauto.json", git_tmp_path)
	err := ioutil.WriteFile(json_path, []byte(json), 0744)
	if err != nil {
		t.Fatal(err)
	}

	err = run(git_tmp_path)
	if err != nil {
		t.Fatal(err)
	}

	deleteLocalRepo(t)
}

func TestRunFileNotFound(t *testing.T) {
	initRepo(t)
	hook(git_tmp_path)

	err := run(git_tmp_path)
	if err == nil {
		t.Fatal(err)
	}

	deleteLocalRepo(t)
}
