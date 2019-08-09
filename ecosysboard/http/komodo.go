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

package http

import (
	"encoding/json"
	"fmt"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/config"
	"github.com/KomodoPlatform/komodo-ecosysboard/ecosysboard/komodo_cache"
	"github.com/google/go-cmp/cmp"
	"github.com/kpango/glg"
	"github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
	"net/http"
	"sort"
	"strings"
	"sync"
)

type CoinInfos struct {
	Ticker                CoinpaprikaTickerData `json:"ticker"`
	BlockLastHash         string                `json:"block_last_hash"`
	BlockInfo             StatusInfo            `json:"status"`
	NodeIsOnline          bool                  `json:"node_is_online"`
	NodeIsSynced          bool                  `json:"node_is_synced"`
	NotarizedHash         string                `json:"notarizedhash"`
	NotarizedTransactions []string              `json:"notarizedtxid"`
	Supply                float64               `json:"supply"`
	CoingeckoData         *CoingeckoCoinData    `json:"additional_data"`
	KomodoCoinID          string                `json:"komodo_coin_id"`
}

func getInfoAboutSpecificCoin(key string, coinpaprikaID string, coingeckoID string) CoinInfos {
	currentCoin := new(CoinInfos)
	if x, found := komodo_cache.GCache.Get(key); found {
		_ = glg.Infof("Retrieve from cache [coin: %s][coinpaprika_id: %s][coingeckoID: %s]", key, coinpaprikaID, coingeckoID)
		currentCoin = x.(*CoinInfos)
	} else {
		_ = glg.Infof("Retrieve realTime [coin: %s][coinpaprika_id: %s][coingeckoID: %s]", key, coinpaprikaID, coingeckoID)
		currentCoin = GetRealTimeCoinInfos(key, coinpaprikaID, coingeckoID)
	}
	return *currentCoin
}

func GetRealTimeCoinInfos(key string, coinpaprikaID string, coingeckoID string) *CoinInfos {
	currentCoin := new(CoinInfos)
	currentCoin.KomodoCoinID = key
	//! Ticker
	var res *CoinpaprikaTickerData
	if strings.Contains(coinpaprikaID, "-") {
		res = CTickerCoinpaprika(coinpaprikaID)
	} else {
		res = new(CoinpaprikaTickerData)
	}
	if coinpaprikaID == "test coin" || res.Symbol == "" {
		res.Symbol = strings.ToUpper(key)
	}
	//! Additional data
	if len(coingeckoID) > 0 && !strings.Contains(coingeckoID, "test coin") {
		currentCoin.CoingeckoData = CCoinsCoingeckoInformation(coingeckoID)
	}
	//! Last block hash
	supply, err, status := CGetSupplyDexstats(key)
	if err != nil || status != http.StatusOK {
		currentCoin.Supply = 0.0
	} else {
		currentCoin.Supply = supply
	}
	currentCoin.BlockLastHash = CDiagnosticInfoFromNodeDexstats("getLastBlockHash", key).LastBlockHash.Lastblockhash
	currentCoin.BlockInfo = CDiagnosticInfoFromNodeDexstats("getInfo", key).Infos
	node := CNodeSyncStatusDexstats(key)
	currentCoin.NodeIsSynced = node.Status == "finished" && node.BlockChainHeight == currentCoin.BlockInfo.Info.Blocks
	currentCoin.NodeIsOnline = currentCoin.BlockInfo.Info.Connections > 2
	if currentCoin.NodeIsSynced && currentCoin.NodeIsOnline {
		currentCoin.NotarizedHash = CBlockHashFromHeightDexstats(key, fmt.Sprintf("%d", currentCoin.BlockInfo.Info.Notarized)).BlockHash
		currentCoin.NotarizedTransactions = CBlockDetailsDexstats(key, currentCoin.NotarizedHash).Tx
	}
	currentCoin.Ticker = *res
	komodo_cache.GCache.Set(key, currentCoin, cache.NoExpiration)
	return currentCoin
}

func GetInformationForSpecificCoinKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	coinName := ctx.UserValue("coin")
	idx := sort.Search(len(config.GConfig.Coins), func(i int) bool { return config.GConfig.Coins[i].Coin >= coinName.(string) })
	if config.GConfig.Coins[idx].Coin != coinName {
		_, _ = ctx.WriteString(`{"error": "This coin does not seem to be part of the komodo ecosystem"}`)
		ctx.SetStatusCode(http.StatusNotFound)
		return
	}
	_ = glg.Infof("find needle: %v", config.GConfig.Coins[idx])
	coinInfo := getInfoAboutSpecificCoin(config.GConfig.Coins[idx].Coin, config.GConfig.Coins[idx].CoinPaprikaID, config.GConfig.Coins[idx].CoinGeckoID)
	if cmp.Equal(CoinInfos{}, coinInfo) {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(200)
	coinInfoJson, _ := json.Marshal(coinInfo)
	ctx.SetBodyString(string(coinInfoJson))
}

func AllInformationsKomodoEcosystem(ctx *fasthttp.RequestCtx) {
	coinInfos := make([]CoinInfos, 0, len(config.GConfig.Coins))
	mutex := sync.RWMutex{}
	var wg sync.WaitGroup
	wg.Add(len(config.GConfig.Coins))
	for _, value := range config.GConfig.Coins {
		go func(key string, coinpaprikaID string, coingeckoID string) {
			defer wg.Done()
			currentCoin := getInfoAboutSpecificCoin(key, coinpaprikaID, coingeckoID)
			mutex.Lock()
			coinInfos = append(coinInfos, currentCoin)
			mutex.Unlock()
		}(value.Coin, value.CoinPaprikaID, value.CoinGeckoID)
	}
	wg.Wait()
	if len(coinInfos) == 0 {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}
	_ = glg.Debug("coinInfos komodo: %v", coinInfos)
	ctx.SetStatusCode(200)
	jsonTicker, _ := json.Marshal(coinInfos)
	ctx.SetBodyString(string(jsonTicker))
}
