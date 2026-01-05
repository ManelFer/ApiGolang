// rota para o endpoint /hero
package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"api-go/handler"

	"github.com/gorilla/mux"
)

func HeroRouter(r *mux.Router, db *sql.DB) {
	// registrar rota /hero
	r.HandleFunc("/hero", handler.GetHeroHandler(db)).Methods(http.MethodGet)
	fmt.Printf("Registered route: GET /hero")
}
