// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"log"
	"os"
	"todos/pkg/infra"
)

func main() {

	infra.InitRouter()
	/*
		log.Print("starting server...")
		http.HandleFunc("/", handler)

		// Determine port for HTTP service.

		port := os.Getenv("PORT")
		if port == "" {
			port = "8081"
			log.Printf("defaulting to port %s", port)
		}

		// Start HTTP server.
		log.Printf("listening on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}

	*/
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

}

/*
func handler(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
*/
