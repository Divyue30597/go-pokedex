package main

import (
	"fmt"

	"github.com/Divyue30597/pokedex/internal/pokeapi"
)

func (p *pokemonCaught) Init() {
	*p = make(map[string]pokeapi.Pokemon)
}

func (p *pokemonCaught) Add(pokemon string, value pokeapi.Pokemon) {
	if *p == nil {
		p.Init()
	}

	if _, ok := (*p)[pokemon]; !ok {
		(*p)[pokemon] = value
	}
}

func (p *pokemonCaught) Get(pokemon string) (pokeapi.Pokemon, bool) {
	if *p == nil {
		return pokeapi.Pokemon{}, false
	}
	val, ok := (*p)[pokemon]

	return val, ok
}

func (p *pokemonCaught) GetAll() ([]string, error) {
	var pokemonArr []string
	if len(*p) == 0 {
		return pokemonArr, fmt.Errorf("your pokedex is empty")
	}

	for pokemon := range *p {
		pokemonArr = append(pokemonArr, pokemon)
	}

	return pokemonArr, nil
}
