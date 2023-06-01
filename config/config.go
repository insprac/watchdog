package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Coin struct {
	Name        string `yaml:"name"`
	DisplayName string `yaml:"displayName"`
}

type Token struct {
	Name        string `yaml:"name"`
	DisplayName string `yaml:"displayName"`
}

type Silo struct {
	Name        string  `yaml:"name"`
	Address     string  `yaml:"address"`
	DisplayName string  `yaml:"displayName"`
	Tokens      []Token `yaml:"tokens"`
}

type MovementAlert struct {
	Name   string `yaml:"name"`
	Change string `yaml:"change"` // change can be a number or percentage e.g. 2.5 or 25.6%
}

type ThresholdAlert struct {
	Name   string `yaml:"name"`
	Amount string `yaml:"amount"`
}

type Alerts struct {
	Movement  []MovementAlert  `yaml:"movement"`
	Threshold []ThresholdAlert `yaml:"threshold"`
}

type Config struct {
	DiscordWebhookUrl string `yaml:"discordWebhookUrl"`
	Coins             []Coin `yaml:"coins"`
	Silos             []Silo `yaml:"silos"`
	Alerts            Alerts `yaml:"alerts"`
}

var config Config

func init() {
	filename := GetConfigFile()
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}
}

func GetDiscordWebhookUrl() string {
	return config.DiscordWebhookUrl
}

func GetCoins() []Coin {
	return config.Coins
}

func GetSilos() []Silo {
	return config.Silos
}

func GetAlerts() Alerts {
	return config.Alerts
}

func GetMovementAlerts() []MovementAlert {
	return config.Alerts.Movement
}

func GetThresholdAlerts() []ThresholdAlert {
	return config.Alerts.Threshold
}
