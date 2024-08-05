package app

import (
	"fmt"
	"go-bonds-scanner/internal/config"
	"go-bonds-scanner/internal/services/moex"
	"math"
)

func Run() {
	bondsResponses, _ := moex.GetAllBonds()

	// todo: Вот ту подумать над рефактором, как лучше сделать, чтобы не было O(n^2)
	for i, resp := range bondsResponses {
		count := len(resp.Securities.Data)

		for k := 0; k < count; k++ {
			var bondPrice float64 = 0
			var bondYield float64 = 0

			bondSecurity := bondsResponses[i].Securities.Data[k]
			bondMarket := bondsResponses[i].Marketdata.Data[k]

			bondName := bondSecurity[1].(string)
			secid := bondSecurity[0].(string)

			if bondSecurity[2] != nil {
				bondPrice = bondSecurity[2].(float64)
			}

			if bondMarket[1] != nil {
				bondYield = bondMarket[1].(float64)
			}

			// кол-во оставшихся месяцев
			bondDuration := math.Floor((bondPrice/30)*100) / 100

			isBondFit := bondYield > config.YIELD_MORE &&
				bondPrice > config.PRICE_MORE &&
				bondPrice < config.PRICE_LESS &&
				bondDuration > config.DURATION_MORE &&
				bondDuration < config.DURATION_LESS

			if isBondFit {
				fmt.Printf("%s (%s) - Доходность: %f Дюрация: %f\n", bondName, secid, bondYield, bondDuration)
			}
		}
	}
}
