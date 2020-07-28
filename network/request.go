package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

func SendReq(url string, chanel chan []byte) {
	var response, err = http.Get(url)
	CheckErr(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	chanel <- Page{url, (len(string(body)))}
}
