package ticker

import (
	"time"
	"demo/modules/sync/price"
)

func RunSyncCoin() {
	CoinSync := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-CoinSync.C:
			btcPrice := price.NewSyncBtcPrice()
			btcPrice.Run()
		}
	}
}
