package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/eediallo/real_time_forum/internal/db"
)

func FetchPrivateMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("FetchPrivateMessages handler invoked")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fetch private messages from the database
	rows, err := db.DB.Query("SELECT SenderUsername, ReceiverUsername, Content, CreatedAt FROM PrivateMessages")
	if err != nil {
		http.Error(w, "Failed to fetch messages: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []map[string]string
	for rows.Next() {
		var senderUsername, receiverUsername, content string
		var createdAt time.Time
		if err := rows.Scan(&senderUsername, &receiverUsername, &content, &createdAt); err != nil {
			http.Error(w, "Failed to scan message: "+err.Error(), http.StatusInternalServerError)
			return
		}
		formattedCreatedAt := createdAt.Format("01/02/2006 03:04 PM")
		message := map[string]string{
			"senderUsername":   senderUsername,
			"receiverUsername": receiverUsername,
			"content":          content,
			"createdAt":        formattedCreatedAt,
		}
		messages = append(messages, message)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages: "+err.Error(), http.StatusInternalServerError)
	}
}
