package network

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kypej/marshal/commonutils"
)

func SendReq(url, limit, ofset string, respBody chan []byte) {
	var client = http.Client{
		Timeout: 5 * time.Second,
	}
	var req, err = http.NewRequest(http.MethodGet, url, nil)
	commonutils.CheckErr(err)
	var query = req.URL.Query()
	query.Add("limit", limit)
	query.Add("offset", ofset)
	req.URL.RawQuery = query.Encode()
	var resp, err1 = client.Do(req)
	commonutils.CheckErr(err1)
	defer resp.Body.Close()
	var body, err2 = ioutil.ReadAll(resp.Body)
	commonutils.CheckErr(err2)
	respBody <- body
}
