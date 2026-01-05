package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type HeroHandler struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Profissao string `json:"profissao"`
}

func GetHeroHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, name, profissao FROM hero"
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("DB query error (%s): %v", query, err)
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var hero []HeroHandler

		for rows.Next() {
			var h HeroHandler
			err := rows.Scan(&h.ID, &h.Name, &h.Profissao)
			if err != nil {
				log.Printf("DB scan error: %v", err)
				http.Error(w, "Database scan error", http.StatusInternalServerError)
				return
			}
			hero = append(hero, h)
		}

		if err := rows.Err(); err != nil {
			log.Printf("DB rows iteration error: %v", err)
			http.Error(w, "Database iteration error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(hero); err != nil {
			log.Printf("JSON encode error: %v", err)
		}
	}
}
