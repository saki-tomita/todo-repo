// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"todos/pkg/adapter"
)

func main() {

	adapter.InitRouter()

}
