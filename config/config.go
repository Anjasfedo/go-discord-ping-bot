package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token      string `json:"Token"`
	BotkPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Read config.json")

	file, err := os.ReadFile("../config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token

	BotPrefix = config.BotkPrefix

	return nil
}
