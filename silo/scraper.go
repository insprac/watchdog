package silo

import (
	"encoding/json"
	"os/exec"
)

type SiloToken struct {
	Name              string
	TotalDeposited    float64
	AvailableToBorrow float64
	Utilization       float64
}

func ScrapeSiloTokens(siloAddress string) ([]SiloToken, error) {
	cmd := exec.Command("node", "./scraper/index.js", siloAddress)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tokens []SiloToken
	err = json.Unmarshal(output, &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
