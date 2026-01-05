// rota para o endpoint /about
package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"api-go/handler"

	"github.com/gorilla/mux"
)

func AboutRouter(r *mux.Router, db *sql.DB) {
	// registrar rota /about
	r.HandleFunc("/about", handler.GetAboutHandler(db)).Methods(http.MethodGet)
	fmt.Printf("Registered route: GET /about")
}
