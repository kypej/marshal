package middlewear

import (
	"fmt"
	"net/url"

	"github.com/kypej/marshal/infrastructure/network"
	"github.com/kypej/marshal/commonutils"
)

func GetPockemons(baseUrl, limit, offset string, respBody chan []byte) {
	var siteUrl, err = url.Parse(baseUrl)
	commonutils.CheckErr(err)
	var q = url.Values{}
	q.Add("limit", limit)
	q.Add("offset", offset)
	siteUrl.RawQuery = q.Encode()
	fmt.Println(siteUrl.String())
	network.SendGetReq(siteUrl.String(), respBody)
}
