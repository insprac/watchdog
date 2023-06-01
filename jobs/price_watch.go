package jobs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/insprac/watchdog/alerts"
	"github.com/insprac/watchdog/config"
	"github.com/insprac/watchdog/utils"
	"github.com/tidwall/gjson"
)

func PriceWatchJob() {
	fmt.Println("Price Watch Job running...")
	coins := config.GetCoins()
	for _, coin := range coins {
		checkPrice(coin.Name)
	}
}

func checkPrice(coin string) {
	price, err := getPrice(coin)
	if err != nil {
		fmt.Printf("Error getting price for %s: %v\n", coin, err)
		return
	}

	alerts.CheckAlerts("coins."+coin, price, "$"+utils.FormatNumber(price))
}

func getPrice(coin string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", coin)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("CoinGecko API request failed with status %d, %s", resp.StatusCode, string(body))
	}

	price := gjson.Get(string(body), coin+".usd").String()
	return strconv.ParseFloat(price, 64)
}
