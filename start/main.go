package main

import (
	"encoding/json"
	"github.com/kypej/marshal/commonutils"
	"github.com/kypej/marshal/server"
	"github.com/kypej/marshal/models"
	"github.com/kypej/marshal/network"
)

func main() {
	var chanelBody = make(chan ([]byte))
	go network.SendReq("https://pokeapi.co/api/v2/pokemon?limit=5&offset=1", chanelBody)
	var model models.PokemonName
	var err = json.Unmarshal(<-chanelBody, &model)
	commonutils.CheckErr(err)
	server.SetPokemonModel(model)
	server.UpServer()
}