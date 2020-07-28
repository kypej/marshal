package main

import (
	"log"

	"github.com/kypej/marshal/network"
)

func main() {
	var chanelBody = make(chan([]byte))
	var body = network.SendReq("https://pokeapi.co/api/v2/pokemon?limit=5", chanelBody)
}