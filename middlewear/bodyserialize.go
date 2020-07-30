package middlewear

import (
	"bytes"
	"encoding/json"

	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/models"
)

var model bytes.Buffer

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
