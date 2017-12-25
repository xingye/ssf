package config

import (
	"errors"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

var token string

func init() {
	if os.Getenv("ssf_token") != "" {
		token = os.Getenv("ssf_token")
	}
}

func GetToken() string {
	if token == "" {
		var err error

		validate := func(input string) error {
			if len(input) == 0 {
				return errors.New("invalidate token.")
			}
			return nil
		}

		prompt := promptui.Prompt{Label: "Input token", Validate: validate}
		token, err = prompt.Run()
		if err != nil {
			log.Fatalln("invalid token.")
			//return ""
		}
	}

	return token
}
