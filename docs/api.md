# Rest Api Doc

## search dexstats information

Retrieves an url from an explorer from a block, addresses or transaction

**URL** : `/api/v1/dexstats/:coin/search`

**Method** : `POST`

**Auth required** : No

**Permissions required** : None

**Data constraints**

```json
{
    "input": "[valid input (transaction, block hash, block height, or address]"
}
```

**Data example**

```json
{
    "input": "06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "url_to_redirect": "http://kmd.explorer.dexstats.info/block/06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"
}
```

Curl command: `curl -X POST http://127.0.0.1:8080/api/v1/dexstats/kmd/search -H 'Content-Type: application/json'  -d '{"input": "06d6747a49097830574cf8d33e399d8a8679e457493cd17390a80d0f916287bc"}'`

## get information about all coins

Get all the information about all the coins from the **komodo ecosystem**: ticker, last block, status, sync, notarized hash

**URL** : `/api/v1/tickers/`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Content examples**

```json
[
   {
      "ticker":{
         "id":"kmd-komodo",
         "name":"Komodo",
         "symbol":"KMD",
         "rank":53,
         "circulating_supply":115225475,
         "total_supply":115225475,
         "max_supply":0,
         "beta_value":0.997236,
         "last_updated":"2019-07-26T10:23:35Z",
         "quotes":{
            "USD":{
               "price":1.16555727,
               "volume_24h":2503791.1344734,
               "volume_24h_change_24h":-28.48,
               "market_cap":134301890,
               "market_cap_change_24h":-0.48,
               "percent_change_1h":-0.25,
               "percent_change_12h":-0.64,
               "percent_change_24h":-0.52,
               "percent_change_7d":-3.64,
               "percent_change_30d":-11.33,
               "percent_change_1y":-31.94,
               "ath_price":15.4149,
               "ath_date":"2017-12-21T08:04:00Z",
               "percent_from_price_ath":-92.44
            }
         }
      },
      "block_last_hash":"08c358a9fcbbcd5b9169d4497dfb136fdb07848e4a25685a4fe9ed822ecc57c0",
      "status":{
         "info":{
            "version":2001526,
            "protocolversion":170007,
            "blocks":1459464,
            "timeoffset":0,
            "connections":118,
            "proxy":"",
            "difficulty":186171523.1663218,
            "testnet":false,
            "relayfee":0.000001,
            "errors":"",
            "notarized":1459450,
            "network":"livenet"
         }
      },
      "node_is_online":true,
      "node_is_synced":true,
      "notarizedhash":"0000000042cc6301354c2f93595cb360587aeb8f255c0597e567290922429088",
      "notarizedtxid":[
         "19a6760bb363b81b9bb382a01eb1a220a981bc7599eb3db8317d1ade8ab01032",
         "a6c385e4b619d0ab2ba0a2fc4a7713435931b2e85b8bea086b1d8a27a0607000",
         "ddab897ad5d5b6c5335db59573aec40b0311d8e2d65aee406f346ccb63fb2ab0",
         "020f86c7d7295cd52bde4f6f8e9efa1fb75170bcff310960dd749d7c8b06f597",
         "a8169cb655a77b587e9c875fb2b8f51001631b90fc6cf7db298628ce1916f3fc",
         "bfc9770c25f521d43443149ca33bd38d2f146676ec2ccf78210cc258eedc5ac0",
         "9c6f7b80900ff3e856124ed3371c31dc0c4667579892f2f2bd641b07fba8b19c"
      ],
      "supply":115225679.67
   },
   {
      "ticker":{
         "id":"k64-komodore64",
         "name":"Komodore64",
         "symbol":"K64",
         "rank":2273,
         "circulating_supply":0,
         "total_supply":64000812,
         "max_supply":64000777,
         "beta_value":1.55205,
         "last_updated":"2019-07-26T10:23:56Z",
         "quotes":{
            "USD":{
               "price":0.23472675,
               "volume_24h":4.694535,
               "volume_24h_change_24h":-99.98,
               "market_cap":0,
               "market_cap_change_24h":0,
               "percent_change_1h":-0.22,
               "percent_change_12h":-1.48,
               "percent_change_24h":-19.07,
               "percent_change_7d":21.16,
               "percent_change_30d":-28.25,
               "percent_change_1y":0,
               "ath_price":0.8719188,
               "ath_date":"2019-06-12T06:27:29Z",
               "percent_from_price_ath":-73.08
            }
         }
      },
      "block_last_hash":"0b4f35467f69ef0268798a307036e151714f98fa46a41069a0c0a0e9b86a6016",
      "status":{
         "info":{
            "version":2001526,
            "protocolversion":170007,
            "blocks":152962,
            "timeoffset":0,
            "connections":83,
            "proxy":"",
            "difficulty":760.0841178010235,
            "testnet":false,
            "relayfee":0.000001,
            "errors":"",
            "notarized":152950,
            "network":"livenet"
         }
      },
      "node_is_online":true,
      "node_is_synced":true,
      "notarizedhash":"021d5054f7372bb0e6bf9ff4fe6ab928f9cd5062bd4e168d38d62e0c82f13c27",
      "notarizedtxid":[
         "f3fde5106f964056c4dd1157da4a493a89c61f415e7d81783bfb4f60f807113d",
         "7d04cdadf1b3d1c19424ec011e05a6ccf76d1bd2308b790b5f49fd1bc5b1a2cd"
      ],
      "supply":64000811.878662
   }
]
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers`

## get information about one specific coin

Get all the information about a coin from the **komodo ecosystem**: ticker, last block, status, sync, notarized hash

**URL** : `/api/v1/tickers/:coin`

**Method** : `GET`

**Auth required** : No

**Permissions required** : None

### Success Response

**Code** : `200 OK`

**Content examples**

For a coin with ID `kmd` that exists in the **komodo** ecosystem.

```json
{
   "ticker":{
      "id":"kmd-komodo",
      "name":"Komodo",
      "symbol":"KMD",
      "rank":53,
      "circulating_supply":115225475,
      "total_supply":115225475,
      "max_supply":0,
      "beta_value":0.997236,
      "last_updated":"2019-07-26T10:22:49Z",
      "quotes":{
         "USD":{
            "price":1.16555727,
            "volume_24h":2503791.1344734,
            "volume_24h_change_24h":-28.48,
            "market_cap":134301890,
            "market_cap_change_24h":-0.48,
            "percent_change_1h":-0.25,
            "percent_change_12h":-0.64,
            "percent_change_24h":-0.52,
            "percent_change_7d":-3.64,
            "percent_change_30d":-11.33,
            "percent_change_1y":-31.94,
            "ath_price":15.4149,
            "ath_date":"2017-12-21T08:04:00Z",
            "percent_from_price_ath":-92.44
         }
      }
   },
   "block_last_hash":"023a49278f88fbc59f5bedd07dcc92072d6b6a36750b466f95336a7789fcc871",
   "status":{
      "info":{
         "version":2001526,
         "protocolversion":170007,
         "blocks":1459463,
         "timeoffset":0,
         "connections":118,
         "proxy":"",
         "difficulty":185855967.5858431,
         "testnet":false,
         "relayfee":0.000001,
         "errors":"",
         "notarized":1459450,
         "network":"livenet"
      }
   },
   "node_is_online":true,
   "node_is_synced":true,
   "notarizedhash":"0000000042cc6301354c2f93595cb360587aeb8f255c0597e567290922429088",
   "notarizedtxid":[
      "19a6760bb363b81b9bb382a01eb1a220a981bc7599eb3db8317d1ade8ab01032",
      "a6c385e4b619d0ab2ba0a2fc4a7713435931b2e85b8bea086b1d8a27a0607000",
      "ddab897ad5d5b6c5335db59573aec40b0311d8e2d65aee406f346ccb63fb2ab0",
      "020f86c7d7295cd52bde4f6f8e9efa1fb75170bcff310960dd749d7c8b06f597",
      "a8169cb655a77b587e9c875fb2b8f51001631b90fc6cf7db298628ce1916f3fc",
      "bfc9770c25f521d43443149ca33bd38d2f146676ec2ccf78210cc258eedc5ac0",
      "9c6f7b80900ff3e856124ed3371c31dc0c4667579892f2f2bd641b07fba8b19c"
   ],
   "supply":115225679.67
}
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers/kmd`

### Error Response

For a coin with ID `nonexistent` that doesnt exists in the **komodo** ecosystem.

**Code** : `404 Not found`

**Content examples**

```json
{
 "error": "This coin does not seem to be part of the komodo ecosystem"
}
```

Curl command: `curl http://127.0.0.1:8080/api/v1/tickers/nonexistent`
