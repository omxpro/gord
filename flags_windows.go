package main

import (
	"flag"

	"github.com/cainy-a/gord/config"
)

func init() {
	flag.BoolVar(&config.DisableUTF8, "disable-UTF8", config.DisableUTF8, "Replaces certain characters with question marks in order to avoid broken rendering")
}
