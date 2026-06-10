package bagoette

import (
	"fmt"
	"net/http"
	"strconv"
)

func (b *BagoetteClient) Serve(port int) error {
	b.SetPort(port)
	b.registerAllRoutes()
	b.ShowRoutes()
	return http.ListenAndServe(":" + strconv.Itoa(b.port), b.httpHandler)
}

func (b *BagoetteClient) ShowRoutes() {
	fmt.Println("Registered Routes:")
	for _, route := range *b.routes {
		fmt.Println("- "+ route.pattern)
	}
}

func (b *BagoetteClient) Close() {
	b.httpClient.CloseIdleConnections()
}