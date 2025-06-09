package main

import (
	"fmt"
	"go-file-soap/internal/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/soap", api.UploadJsonHandler)

	fmt.Println("SOAP service listening on http://localhost:8080/soap")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
