package handler

import (
	"encoding/json"
	"log"
	"net/http"

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
		var senderUsername, receiverUsername, content, createdAt string
		if err := rows.Scan(&senderUsername, &receiverUsername, &content, &createdAt); err != nil {
			http.Error(w, "Failed to scan message: "+err.Error(), http.StatusInternalServerError)
			return
		}
		message := map[string]string{
			"senderUsername":   senderUsername,
			"receiverUsername": receiverUsername,
			"content":          content,
			"createdAt":        createdAt,
		}
		messages = append(messages, message)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages: "+err.Error(), http.StatusInternalServerError)
	}
}
