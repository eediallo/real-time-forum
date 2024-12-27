package handler

import (
	"net/http"
	"time"

	"github.com/eediallo/real_time_forum/internal/db"
	"github.com/eediallo/real_time_forum/internal/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		registrationDate := time.Now()

		// Check if the username already exists
		exists, err := checkEmailExists(email)
		if err != nil {
			http.Error(w, "Email already exists in db", http.StatusInternalServerError)
			return
		}
		if exists {
			w.WriteHeader(http.StatusBadRequest)
			ErrorPageHandler(w, "Email already taken", nil)
			return
		}

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.DB.Exec(addUserDetailsQuery, username, email, hashedPassword, registrationDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	} else {
		data := PageData{
			CssRegisterPath: cssRegisterPath,
			HomePath:        homePagePath,
			LogoCSS:         cssLogoPath,
			Logo:            logPath,
		}
		RenderTemplate(w, "register", data)
	}
}

func checkEmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM User WHERE Email=?)"
	err := db.DB.QueryRow(query, email).Scan(&exists)
	return exists, err
}
