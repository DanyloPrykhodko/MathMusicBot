package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/drprykhodko/MathMusicBot/internal/app/bot"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/bot.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := bot.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := bot.Start(config); err != nil {
		log.Fatal(err)
	}
}
