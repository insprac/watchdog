package jobs

import (
	"fmt"

	"github.com/insprac/watchdog/alerts"
	"github.com/insprac/watchdog/config"
	"github.com/insprac/watchdog/silo"
	"github.com/insprac/watchdog/utils"
)

func SiloWatchJob() {
	fmt.Println("Checking silos...")
	silos := config.GetSilos()
	for _, currentSilo := range silos {
		tokens, err := silo.ScrapeSiloTokens(currentSilo.Address)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, siloToken := range currentSilo.Tokens {
			for _, token := range tokens {
				if token.Name == siloToken.DisplayName {
					name := "silos." + currentSilo.Name + "." + siloToken.Name
					alerts.CheckAlerts(name+".totalDeposited", token.TotalDeposited, utils.FormatNumber(token.TotalDeposited))
					alerts.CheckAlerts(name+".availableToBorrow", token.AvailableToBorrow, utils.FormatNumber(token.AvailableToBorrow))
					alerts.CheckAlerts(name+".utilization", token.Utilization, utils.FormatNumber(token.Utilization)+"%")
				}
			}
		}
	}
}
