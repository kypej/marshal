package server

import (
	"bytes"
	"log"
	"net/http"

	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/middlewear"
)

func UpServer() {
	http.HandleFunc("/pokemons", pokemonsHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func write(writer http.ResponseWriter, message bytes.Buffer) {
	_, err := writer.Write(message.Bytes())
	commonutils.CheckErr(err)
}

func pokemonsHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		var body bytes.Buffer
		var _, err = body.Write([]byte(`404 page not found`))
		commonutils.CheckErr(err)
		writer.WriteHeader(404)
		_, err = writer.Write(body.Bytes())
		commonutils.CheckErr(err)
		return
	}
	middlewear.PockemonsHandlers(writer, request)

}
