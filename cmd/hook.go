/*
Copyright © 2019 Abdulaziz Alfoudari <aziz.alfoudari@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
