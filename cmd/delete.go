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
	"ssf/slack"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	deleteCmd.Flags().StringVarP(&user, "user", "u", "", "slack user id")
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete all slack shared files",
	Run: func(cmd *cobra.Command, args []string) {
		if success, fail, err := slack.DeleteAllFiles(user); err != nil {
			log.Error().Msgf("delete all fils error:%+v\n", err)
		} else {
			if len(fail) > 0 {
				log.Info().Msgf("%d files has deleted successfully.\n", len(success))
				log.Warn().Msgf("%d files failed.\n", len(fail))
			} else {
				log.Info().Msgf("Great! Delete %d files successfully.\n", len(success))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
