package main

import (
	"fmt"
	// "log"
	"net/http"
	"todo-project/router"

	"github.com/rs/cors"
)

func main() {
	route := router.Router()
	fmt.Println("Começando na porta 9000")

	// Crie uma instância do CORS com as configurações desejadas
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Adicione aqui as origens permitidas
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	// Use o middleware CORS com o roteador
	handler := c.Handler(route)

	// Inicie o servidor
	http.ListenAndServe(":9000", handler)
}
