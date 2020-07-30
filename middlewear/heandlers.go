package middlewear

import (
	"encoding/json"
	"net/http"

	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/models"
	"github.com/kypej/marshal/varible"
)

func PockemonsHandlers(writer http.ResponseWriter, request *http.Request) {
	var args = request.URL.Query()
	var limit = args.Get("limit")
	var offset = args.Get("offset")
	var chanelBody = make(chan ([]byte))
	go GetPockemons(varible.BaseURL, limit, offset, chanelBody)
	var responseModel models.PokemonName
	var err = json.Unmarshal(<-chanelBody, &responseModel)
	commonutils.CheckErr(err)
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	var body = GenerateJsonBody(responseModel)
	writer.Write(body.Bytes())
}
