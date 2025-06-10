package main

import (
	"fmt"
	"go-file-soap/internal/api"
	"go-file-soap/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file:", err)
	}

	expectedToken := os.Getenv("API_TOKEN")
	if expectedToken == "" {
		log.Fatal("API_TOKEN parameter not found in environment variables file:", err)
	}

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(expectedToken))
		r.Post("/soap", api.UploadJsonHandler)
		r.Get("/soap.wsdl", api.WSDLHandler)
	})

	fmt.Println("SOAP service listening on http://localhost:8080/api/soap")
	log.Fatal(http.ListenAndServe(":8080", r))
}
