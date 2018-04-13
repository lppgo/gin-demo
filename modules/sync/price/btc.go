package price

import (
	"net/url"
)

var Apikey = "xxxx"
var Url = "https://xxxxxxxxxxxx/xxxxxx"

type SyncBtcPrice struct {
	SyncPrice
}

type PriceResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSyncBtcPrice() *SyncBtcPrice {
	SyncBtcPrice := &SyncBtcPrice{}
	return SyncBtcPrice
}

func GetUrl() string {
	return Url
}

func (syncBtcPrice *SyncBtcPrice) getCurrencyParams() url.Values {
	params := url.Values{}
	params.Set("key", Apikey)
	return params
}

func (syncBtcPrice *SyncBtcPrice) Run() {

}
