package main
package internal/router/routers


import (
	"fmt"
	"github.com/Draztick/lokic2/internal/routers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const serverAddr = "127.0.0.1"
const serverPort = "8080"

func main() {
	r := chi.NewRouter()

	router.SetupRoutes(r)

	serverAddrPort := fmt.Sprintf(format: "%s:%s", serverAddr, serverPort)

	log.Printf(format: "Server listening on %s\n", serverAddrPort)

	err := http.ListenAndServe(serverAddrPort, r)
	if err != nil {
		log.Fatal(err)
	}
}
