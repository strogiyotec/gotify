/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

type resource struct {
	Version string
	Path    string
}

func NewResource() resource {
	return resource{
		Version: "-v",
	}
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a list of binaries that can be updated",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		checkVersions(path)
	},
}

func checkVersions(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e)
	}
	var resources []resource
	e = json.Unmarshal([]byte(file), &resources)
	if e != nil {
		panic(e)
	}
	fmt.Println(resources[0].Version)

}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("path", "p", os.Getenv("XDG_CONFIG_HOME")+"/gotify/binaries.json", "Set config path")
}
