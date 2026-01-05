// ping
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"api-go/database"
	"api-go/router"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	// tentativa direta (padrão)
	if err := godotenv.Load(); err == nil {
		return nil
	}
	// procurar .env nos diretórios pais (máx 6 níveis)
	wd, _ := os.Getwd()
	curr := wd
	for i := 0; i < 6; i++ {
		candidate := filepath.Join(curr, ".env")
		if _, err := os.Stat(candidate); err == nil {
			if err := godotenv.Load(candidate); err == nil {
				log.Printf("Loaded .env from %s", candidate)
				return nil
			}
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return fmt.Errorf("could not find .env file in working dir or parent dirs")
}

func main() {
	// carregar o .env (procura no diretório atual e nos pais)
	if err := loadEnv(); err != nil {
		log.Printf("Não foi possível carregar .env: %v", err)
	}

	// conectar no banco
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods("GET")
	// rota /about usando router que registra o handler com acesso ao DB
	router.AboutRouter(r, db)
	router.HeroRouter(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Servidor rodando em Port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
