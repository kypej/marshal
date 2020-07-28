package network

import (
	"io/ioutil"
	"net/http"

	"github.com/kypej/marshal/commonutils"
)

func SendReq(url string, respBody chan []byte) {
	var resp, err = http.Get(url)
	commonutils.CheckErr(err)
	defer resp.Body.Close()
	var body, err1 = ioutil.ReadAll(resp.Body)
	commonutils.CheckErr(err1)
	respBody <- body
}
