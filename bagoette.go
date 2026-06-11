package bagoette

import (
	"log"
	"os"
)

const version = "v0.2.3 (Early Development)"
const author = "Sevalino Elfata"
const repo = "https://github.com/sevelfatt/bagoette"

var logger = log.New(os.Stdout, Red + "Bagoette: " + Reset, log.LstdFlags)