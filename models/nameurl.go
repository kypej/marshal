package models

type PokemonName struct{
	Count int `json:"count"`
    Next string
	Previous string
	Results []nameUrl
}

type nameUrl struct {
	Name string
	Url string
} 