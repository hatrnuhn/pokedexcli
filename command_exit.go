package main

import "os"

func commandExit(cfg *config, _ string, _ map[string]bool) error {
	os.Exit(0)
	return nil
}
