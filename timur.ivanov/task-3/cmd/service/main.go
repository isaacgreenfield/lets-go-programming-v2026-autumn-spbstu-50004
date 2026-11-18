package main

import (
	"flag"
	"fmt"

	"task-3/internal/config"
	"task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("no path")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	err = currency.Process(cfg.InputFile, cfg.OutputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
