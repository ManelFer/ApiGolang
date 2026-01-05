package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type AboutHandler struct {
	ID             int64  `json:"id"`
	TitleAbout     string `json:"title_about"`
	ParagraphAbout string `json:"paragraph_about"`
}

func GetAboutHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, title_about, paragraph_about FROM about"
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("DB query error (%s): %v", query, err)
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var about []AboutHandler

		for rows.Next() {
			var a AboutHandler
			err := rows.Scan(&a.ID, &a.TitleAbout, &a.ParagraphAbout)
			if err != nil {
				log.Printf("DB scan error: %v", err)
				http.Error(w, "Database scan error", http.StatusInternalServerError)
				return
			}
			about = append(about, a)
		}

		if err := rows.Err(); err != nil {
			log.Printf("DB rows iteration error: %v", err)
			http.Error(w, "Database iteration error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		if err := json.NewEncoder(w).Encode(about); err != nil {
			log.Printf("JSON encode error: %v", err)
		}
	}
}
