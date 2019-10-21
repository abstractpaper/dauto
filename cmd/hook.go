package cmd

import (
	"fmt"
	"io/ioutil"
	"path"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Hook a git repository",
	Long:  `Write a pre-commit git hook into a repository to enable dauto integration.`,
	Args:  cobra.ExactArgs(1),
	Run:   Hook,
}

func Hook(cmd *cobra.Command, args []string) {
	hook(args...)
}

func hook(args ...string) {
	repo_path := path.Clean(args[0])
	sh_path := fmt.Sprintf("%s/.git/hooks/pre-commit", repo_path)

	// write a pre-commit executable shell script to run dauto
	sh := `#!/bin/bash
	dauto run %s`
	sh = fmt.Sprintf(sh, repo_path)

	err := ioutil.WriteFile(sh_path, []byte(sh), 0744)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Created pre-commit hook in %s", sh_path))
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
