package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/models"
	"github.com/kypej/marshal/network"
	"github.com/kypej/marshal/varible"
)

var model bytes.Buffer

func UpServer() {
	http.HandleFunc("/pokemon", handler)
	http.HandleFunc("/custom", customHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// func viewHandler(writer http.ResponseWriter, request *http.Request) {
// 	message := []byte("Hello, web!")
// 	_, err := writer.Write(message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func write(writer http.ResponseWriter, message bytes.Buffer) {
	_, err := writer.Write(message.Bytes())
	commonutils.CheckErr(err)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	write(writer, model)
}

func customHandler(writer http.ResponseWriter, request *http.Request) {
	var err = request.ParseForm()
	commonutils.CheckErr(err)
	// var argsMap = make(map[string][]string)
	// for k, v := range request.Form {
		// argsMap[k] = v
	// }
	// var limit = argsMap["limit"][0]
	// var ofset = argsMap["offset"][0]
	var args = request.URL.Query()
	var limit = args.Get("limit")
	var ofset = args.Get("offset")
	var chanelBody = make(chan ([]byte))
	go network.SendReq(varible.BaseURL, limit, ofset, chanelBody)
	var responseModel models.PokemonName
	err = json.Unmarshal(<-chanelBody, &responseModel)
	commonutils.CheckErr(err)
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	write(writer, GenerateJsonBody(responseModel))
}

func GenerateJsonBody(model models.PokemonName) bytes.Buffer {
	var marshalBody, err = json.Marshal(model.Results)
	commonutils.CheckErr(err)
	var body bytes.Buffer
	_, err = body.Write(marshalBody)
	commonutils.CheckErr(err)
	return body
}

func SetPokemonModel(pokemonModel models.PokemonName) {
	model.Reset()
	model.Grow(512)
	for _, element := range pokemonModel.Results {
		model.Write([]byte(element.Name))
		model.Write([]byte(" "))
		model.Write([]byte(element.URL))
		model.Write([]byte("\n"))
	}
}
