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
	"log"
	"ssf/model"
	"ssf/slack"
	"time"

	"github.com/spf13/cobra"
)

var before string
var after string
var fileIds []string
var days int

func init() {
	deleteCmd.Flags().StringVarP(&user, "user", "u", "", "slack user id")
	deleteCmd.Flags().IntVarP(&days, "days", "d", 0, "delete all files that were created more than 'days' late")
	deleteCmd.Flags().StringVarP(&before, "before", "b", "", "delete all files that were created before the date(YYYY-MM-DD)")
	deleteCmd.Flags().StringVarP(&after, "after", "a", "", "delete all files that were created after the date(YYYY-MM-DD)")
	deleteCmd.Flags().StringSliceVarP(&fileIds, "files", "f", nil, "delete the specific files(F7F9FQM50,F7GC0DL9M...)")
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete all slack shared files",
	Run: func(cmd *cobra.Command, args []string) {
		flagSet := cmd.Flags()
		if (flagSet.NFlag() == 1 && user != "") || (flagSet.NFlag() == 0) {
			deleteAll()
			return
		}

		if len(fileIds) > 0 {
			delete(fileIds)
			return
		}

		var filter func(model.File) bool

		if before != "" {
			filter = beforeFilter
		} else if after != "" {
			filter = afterFilter
		} else if days > 0 {
			filter = daysFilter
		}

		success, fail, err := slack.DeleteFilesWithFilter(user, filter)
		if err != nil {
			log.Println("delete fail. error:", err)
			return
		}
		logSummary(success, fail)
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

func deleteAll() {
	if success, fail, err := slack.DeleteAllFiles(user); err != nil {
		log.Fatalf("delete all fils error:%+v\n", err)
	} else {
		if len(fail) > 0 {
			log.Printf("%d files have deleted successfully.\n", len(success))
			log.Printf("%d files failed.\n", len(fail))
		} else {
			log.Printf("Great! Delete %d files successfully.\n", len(success))
		}
	}
}

func delete(files []string) {
	logSummary(slack.DeleteFiles(files))
}

func parse(date string) (*time.Time, error) {
	layout := "2006-01-02"
	result, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func validate(cmd *cobra.Command) {
	flagSet := cmd.Flags()
	count := 1
	if flagSet.Lookup("user") != nil {
		count = 2
	}

	if flagSet.NFlag() > count {
		log.Fatalln("a, b, d, f option must be exclusive")
	}
}

func beforeFilter(file model.File) bool {
	date, err := parse(before)
	if err != nil {
		log.Fatalf("parse daete error:%+v\nPlease use YYYY-MM-DD format", err)
	}
	if file.CreatedDateWithoutTime().Before(*date) {
		return true
	}
	return false
}

func afterFilter(file model.File) bool {
	date, err := parse(after)
	if err != nil {
		log.Fatalf("parse daete error:%+v\nPlease use YYYY-MM-DD format", err)
	}
	if file.CreatedDateWithoutTime().After(*date) {
		return true
	}
	return false
}

func daysFilter(file model.File) bool {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	if file.CreatedDateWithoutTime().Sub(now).Hours()/24 > float64(days) {
		return true
	}
	return false
}

func logSummary(success []string, fail []string) {
	fmt.Printf("%d files have deleted successfully.\n", len(success))
	fmt.Printf("%d files failed.\n", len(fail))
}
