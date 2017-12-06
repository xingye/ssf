// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"ssf/slack"
	"text/tabwriter"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var user string

func init() {
	listCmd.Flags().StringVarP(&user, "user", "u", "", "slack user id")
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all slack shared files",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := slack.ListAllFiles(user)
		if err != nil {
			log.Error().Msgf("list files error:%+v\n", err)
			return
		}

		const nameLen = 20
		const layout = "2006-01-02 15:04-05"
		var date time.Time
		var dateStr string

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
		fmt.Fprintln(w, "Id\tName\tType\tSize\tCreate\t")

		for _, file := range files {
			name := file.Name
			if len(name) > nameLen {
				name = "..." + name[len(name)-nameLen:]
			}

			date = time.Unix(file.Created, 0)
			dateStr = date.Format(layout)

			fmt.Fprintf(w, "%s\t%s\t%s\t%.2fKB\t%s\t\n", file.Id, name, file.Type, file.Size/1024, dateStr)
		}
		w.Flush()

		fmt.Printf("\nTotal files: %d\n", len(files))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
