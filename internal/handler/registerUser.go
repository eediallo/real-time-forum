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
		nickName := r.FormValue("nickName")
		age := r.FormValue("age")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		gender := r.FormValue("gender")
		email := r.FormValue("email")
		password := r.FormValue("password")
		registrationDate := time.Now()

		// Check if the username already exists
		exists, err := checkEmailExistsOrNickName(email, nickName)
		if err != nil {
			http.Error(w, "Emai/NicKName already exists in db", http.StatusInternalServerError)
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

		_, err = db.DB.Exec(addUserDetailsQuery, nickName, age, firstName, lastName, gender, username, email, hashedPassword, registrationDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/users/login", http.StatusSeeOther)
	} else {
		data := db.PageData{
			CssRegisterPath:    cssRegisterPath,
			HomePath:           homePagePath,
			LogoCSS:            cssLogoPath,
			Logo:               logPath,
			SignUpCoverImage:   signUpCoverImagePath,
			ValidatePasswordJS: validatePasswordJS,
		}
		RenderTemplate(w, "register", data)
	}
}

func checkEmailExistsOrNickName(email, nickName string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM User WHERE Email=? OR NickName=?)"
	err := db.DB.QueryRow(query, email, nickName).Scan(&exists)
	return exists, err
}
