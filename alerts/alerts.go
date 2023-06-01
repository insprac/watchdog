package alerts

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/insprac/watchdog/config"
	"github.com/insprac/watchdog/discord"
	"github.com/insprac/watchdog/state"
	"github.com/insprac/watchdog/utils"
)

func CheckAlerts(name string, value float64, displayValue string) {
	last := state.Get(name)

	if last == nil {
		state.Set(name, value)
		message := fmt.Sprintf("Added %s: %s", name, displayValue)
		Alert(message)
		return
	}

	alerts := config.GetAlerts()

	for _, alert := range alerts.Movement {
		if alert.Name == name {
			if checkMovementAlert(alert, value, last.Value, displayValue) {
				state.Set(name, value)
			}
		}
	}

	for _, alert := range alerts.Threshold {
		if alert.Name == name {
			if checkThresholdAlert(alert, value, last.Value, displayValue) {
				state.Set(name, value)
			}
		}
	}
}

func checkMovementAlert(alert config.MovementAlert, value float64, lastValue float64, displayValue string) bool {
	changeStr := alert.Change
	isPercentage := regexp.MustCompile(`%$`).MatchString(changeStr)
	change := stringToFloat(strings.TrimSuffix(changeStr, "%"))

	if isPercentage {
		percentageChange := ((value - lastValue) / lastValue) * 100
		if percentageChange >= change || percentageChange <= -change {
			message := fmt.Sprintf("%s has moved by %.2f%% and is now %s", alert.Name, percentageChange, displayValue)
			Alert(message)
			return true
		}
	} else {
		if value-change >= lastValue || value+change <= lastValue {
			message := fmt.Sprintf("%s has moved by %s and is now %s", alert.Name, utils.FormatNumber(value-lastValue), displayValue)
			Alert(message)
			return true
		}
	}

	return false
}

func checkThresholdAlert(alert config.ThresholdAlert, value float64, lastValue float64, displayValue string) bool {
	threshold := stringToFloat(alert.Amount)

	if (lastValue < threshold && value >= threshold) || (lastValue > threshold && value <= threshold) {
		message := fmt.Sprintf("%s has crossed the threshold of %s and is now %s\n", alert.Name, alert.Amount, displayValue)
		Alert(message)
		return true
	}

	return false
}

func Alert(message string) {
	fmt.Println(message)
	discord.SendWebhook(message)
}

func stringToFloat(input string) float64 {
	input = strings.TrimSpace(input)

	length := len(input)
	lastChar := input[length-1:]
	valueStr := input[:length-1]
	multiplier := 1.0

	switch strings.ToLower(lastChar) {
	case "k":
		multiplier = 1e3
	case "m":
		multiplier = 1e6
	case "b":
		multiplier = 1e9
	default:
		valueStr = input
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Printf("Error parsing string to float: %v\n", err)
		return 0
	}

	return value * multiplier
}
