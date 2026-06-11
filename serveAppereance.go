package bagoette

import "fmt"

var banner string = `
 ____                         _   _       
| __ )  __ _  __ _  ___   ___| |_| |_ ___ 
|  _ \ / _` + "`" + ` |/ _` + "`" + ` |/ _ \ / _ \ __| __/ _ \
| |_) | (_| | (_| | (_) |  __/ |_| ||  __/\
|____/ \__,_|\__, |\___/ \___|\__|\__\___|
             |___/                        
`

var version string = "v0.2.3 (Early Development)"

func (b *BagoetteClient) ServeAppearance() {
	fmt.Println(banner)
	fmt.Println("Version: ", version)
	fmt.Println("By: ", "Sevalino Elfata\n")

	b.ShowRoutes()
	fmt.Println("\nServer is running on port", b.port)

}

func (b *BagoetteClient) ShowRoutes() {
	fmt.Println("Registered Routes:")
	for _, route := range *b.routes {
		fmt.Println("- "+ route.pattern)
	}
}
