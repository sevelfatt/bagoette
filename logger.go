package bagoette

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, Red + "[BAGOETTE] " + Reset, log.LstdFlags)