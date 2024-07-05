package main

import "os"

func commandExit(cfg *config, location string) error {
	os.Exit(0)
	return nil
}
