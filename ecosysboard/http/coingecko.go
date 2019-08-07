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
	"github.com/kpango/glg"
	"github.com/valyala/fasthttp"
)

type CoingeckoCoinData struct {
	Links struct {
		Homepage                    []string `json:"homepage"`
		BlockchainSite              []string `json:"blockchain_site"`
		OfficialForumURL            []string `json:"official_forum_url"`
		ChatURL                     []string `json:"chat_url"`
		AnnouncementURL             []string `json:"announcement_url"`
		TwitterScreenName           string   `json:"twitter_screen_name"`
		FacebookUsername            string   `json:"facebook_username"`
		BitcointalkThreadIdentifier int      `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string   `json:"telegram_channel_identifier"`
		SubredditURL                string   `json:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `json:"github"`
			Bitbucket []interface{} `json:"bitbucket"`
		} `json:"repos_url"`
	} `json:"links"`
}

func CCoinsCoingeckoInformation(coinsId string) *CoingeckoCoinData {
	coinsInfo := new(CoingeckoCoinData)
	finalEndpoint := CoingGeckoEndpoint + "/coins/" + coinsId
	req, res := InternalExecGet(finalEndpoint, nil, false)
	if res.StatusCode() == 200 {
		_ = json.Unmarshal(res.Body(), &coinsInfo)
	} else if res.StatusCode() == 429 {
		_ = glg.Warnf("To much request, please retry in one seconds")
	}
	ReleaseInternalExecGet(req, res)
	return coinsInfo
}

func PingCoingecko(ctx *fasthttp.RequestCtx) {
	finalEndpoint := CoingGeckoEndpoint + "/ping"
	InternalExecGet(finalEndpoint, ctx, true)
}
