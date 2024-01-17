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

func ReadConfig() {

}
