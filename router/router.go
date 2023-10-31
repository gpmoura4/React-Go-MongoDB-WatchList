package router

import (
	"todo-project/middleware"

	"github.com/gorilla/mux"
)

// CRIANDO ROTAS
func Router() *mux.Router {

	router := mux.NewRouter()
	middleware.start()
	// router.HandleFunc("api/", middleware.loadEnv).Methods("GET, OPTIONS")
	// Listar
	router.HandleFunc("api/anime", middleware.GetAllAnimes).Methods("GET, OPTIONS")
	// Criar anime route
	router.HandleFunc("api/animes", middleware.CreateAnime).Methods("POST, OPTIONS")
	// Completar anime
	router.HandleFunc("api/animes/{id}", middleware.AnimeFinished).Methods("PUT, OPTIONS")
	// Descompletar anime
	router.HandleFunc("api/descompletarAnime/{id}", middleware.UndoAnime).Methods("PUT, OPTIONS")
	// Apagar anime
	router.HandleFunc("api/apagarAnime/{id}", middleware.DeleteAnime).Methods("DELETE, OPTIONS")
	// Apagar todos os animes
	router.HandleFunc("api/apagarTudo/{id}", middleware.DeleteAllAnimes).Methods("DELETE, OPTIONS")
	return router
}
