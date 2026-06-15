package bagoette

import (
	"log"
	"os"
)

const version = "v0.4.3 (Development)"
const author = "Sevalino Elfata"
const repo = "https://github.com/sevelfatt/bagoette"

var logger = log.New(os.Stdout, Red + "[BAGOETTE] " + Reset, log.LstdFlags)