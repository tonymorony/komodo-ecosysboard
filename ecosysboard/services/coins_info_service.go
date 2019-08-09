/******************************************************************************
 * Copyright © 2013-2019 The Komodo Platform Developers.                      *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * Komodo Platform software, including this file may be copied, modified,     *
 * propagated or distributed except according to the terms contained in the   *
 * LICENSE file                                                               *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/

package services

import (
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/http"
	"github.com/kpango/glg"
	"sync"
	"time"
)

func fetchCoinInfos() {
	_ = glg.Infof("fetch coin info begin")
	var wg sync.WaitGroup
	wg.Add(len(config.GConfig.Coins))
	for _, value := range config.GConfig.Coins {
		go func(key string, coinpaprikaID string, coingeckoID string) {
			defer wg.Done()
			_ = http.GetRealTimeCoinInfos(key, coinpaprikaID, coingeckoID)
		}(value.Coin, value.CoinPaprikaID, value.CoinGeckoID)
	}
	wg.Wait()
	_ = glg.Infof("fetch coin info finish")
}

func LaunchCoinsInfoService() {
	fetchCoinInfos()
	gitStatsTicker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-gitStatsTicker.C:
			fetchCoinInfos()
		}
	}
}
