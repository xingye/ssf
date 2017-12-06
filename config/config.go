package config

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
)

type SlackInfo struct {
	Token string
	User  string
}

var Slack = SlackInfo{}

func Initialize() {
	var err error
	token := os.Getenv("ssf_token")
	if token == "" {
		validate := func(input string) error {
			if len(input) == 0 {
				return errors.New("invalidate token.")
			}
			return nil
		}

		prompt := promptui.Prompt{Label: "Token", Validate: validate}
		token, err = prompt.Run()
		if err != nil {
			log.Fatal().Msgf("invalid token. error:%+v\n", err)
			return
		}
	}
	Slack.Token = token

	if os.Getenv("ssf_user") != "" {
		Slack.User = os.Getenv("ssf_user")
	}
}
