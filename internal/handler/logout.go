package handler

import (
	"net/http"

	"github.com/eediallo/real_time_forum/internal/db"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
		return
	}

	_, err = db.DB.Exec("DELETE FROM Session WHERE SessionID = ?", cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie = &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}
