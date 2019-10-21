package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a job",
	Long:  `Run a job`,
	Args:  cobra.ExactArgs(1),
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	run(args...)
}

func run(args ...string) error {
	log.Printf("%s", Green("Dauto: starting..."))

	// validate args
	dir := args[0]
	fi, err := os.Stat(dir)
	if err != nil {
		log.Fatal(err)
	}
	dautofilePath := ""
	if fi.Mode().IsDir() {
		dautofilePath = filepath.Clean(fmt.Sprintf("%s/dauto.json", dir))
	}

	// parse dauto.json
	dautoFile, err := os.Open(dautofilePath)
	if err != nil {
		log.Error(err)
		return err
	}
	defer dautoFile.Close()

	bytes, _ := ioutil.ReadAll(dautoFile)

	var config interface{}
	json.Unmarshal([]byte(bytes), &config)

	// go over commands and execute them
	for _, element := range config.([]interface{}) {
		log.Printf("%s %s\n", Bold("Running:"), element)
		out, err := exec.Command("sh", "-c", element.(string)).Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("> %s", out)
	}

	log.Printf("%s", Green("Dauto: successful!"))

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
