package bagoette

import (
	"fmt"
	"strconv"
)

const (
    Reset  = "\033[0m"
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Blue   = "\033[34m"
	White = "\033[37m"
)

var RequestStatusColours = map[int]string{
	1: Blue,
	2: Green,
	3: Blue,
	4: Yellow,
	5: Red,
}

var methodColors = map[string]string{
	"GET": Green,
	"POST": Yellow,
	"PUT": Blue,
	"DELETE": Red,
}

var banner string = `
 ____                         _   _       
| __ )  __ _  __ _  ___   ___| |_| |_ ___ 
|  _ \ / _` + "`" + ` |/ _` + "`" + ` |/ _ \ / _ \ __| __/ _ \
| |_) | (_| | (_| | (_) |  __/ |_| ||  __/
|____/ \__,_|\__, |\___/ \___|\__|\__\___|
             |___/                        
`

func (b *BagoetteClient) ServeAppearance() {
	fmt.Println(Red + banner + Reset)
	fmt.Println(Blue + repo + Reset)
	fmt.Println(Red + "\nVersion: " + Reset, version)
	fmt.Println(Red + "By: " + Reset, author+ "\n")

	b.ShowRoutes()
	fmt.Println(Red + "\nServer is running on port", b.opts.Port, Reset)

}

func (b *BagoetteClient) ShowRoutes() {
	fmt.Println("Registered Routes:")
	for _, route := range *b.routes {
		fmt.Println("- ["+methodColors[route.method]+route.method+Reset+"] "+route.path)
	}
}

func Log(message string) {
	logger.Println(message)
}

func getRequestStatusColour(status int) string {
	return RequestStatusColours[status/100]
}

func requestLog(method string, path string, status int) {
	logger.Println(methodColors[method] + method + Reset + " " + path + " " + getRequestStatusColour(status) + strconv.Itoa(status) + Reset)
}