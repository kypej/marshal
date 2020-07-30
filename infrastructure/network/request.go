package network

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kypej/marshal/commonutils"
)

func SendGetReq(url string, respBody chan []byte) {
	var client = http.Client{
		Timeout: 5 * time.Second,
	}
	var req, err = http.NewRequest(http.MethodGet, url, nil)
	commonutils.CheckErr(err)
	var resp, err1 = client.Do(req)
	commonutils.CheckErr(err1)
	defer resp.Body.Close()
	var body, err2 = ioutil.ReadAll(resp.Body)
	commonutils.CheckErr(err2)
	respBody <- body
}

