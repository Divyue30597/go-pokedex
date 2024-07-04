package main

import (
	"errors"
	"fmt"
)

func commandGetMap(cfg *config) error {
	response, err := cfg.pokeapi.GetMapLists(cfg.nextLocaleString)
	if err != nil {
		return err
	}

	cfg.nextLocaleString = response.Next
	cfg.prevLocalString = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandGetPreviousMap(cfg *config) error {
	if cfg.prevLocalString == nil {
		return errors.New("no previous page found or it is the first page")
	}

	response, err := cfg.pokeapi.GetMapLists(cfg.prevLocalString)
	if err != nil {
		return err
	}

	cfg.nextLocaleString = response.Next
	cfg.prevLocalString = response.Previous

	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}
