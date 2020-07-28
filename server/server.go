package server

import (
	"bytes"
	"log"
	"net/http"

	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/models"
)

var model bytes.Buffer

func UpServer() {
	http.HandleFunc("/pokemon", handler)
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

func SetPokemonModel(pokemonModel models.PokemonName) {
	model.Grow(512)
	for _, element := range pokemonModel.Results {
		model.Write([]byte(element.Name))
		model.Write([]byte(" "))
		model.Write([]byte(element.Url))
		model.Write([]byte("\n"))
	}
}
