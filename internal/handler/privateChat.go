package handler

import (
    "log"
    "net/http"

    "github.com/eediallo/real_time_forum/internal/db"
)

func PrivateChat(w http.ResponseWriter, r *http.Request) {
    log.Println("PrivateChat handler invoked")

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Failed to parse form data", http.StatusBadRequest)
        return
    }

    // Get chat input value
    chatContent := r.FormValue("chatInput")
    if chatContent == "" {
        http.Error(w, "Message cannot be empty", http.StatusBadRequest)
        return
    }

    // Insert chat content into the database
    _, err = db.DB.Exec("INSERT INTO PrivateMessages (Content) VALUES (?)", chatContent)
    if err != nil {
        http.Error(w, "Failed to store message: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Message successfully sent."))
    log.Println("Message successfully sent")
}