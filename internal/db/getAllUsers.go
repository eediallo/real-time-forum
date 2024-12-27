package db

import "time"

// User represents a user in the system
type User struct {
    UserID           int       `json:"user_id"`
    NickName         string    `json:"nickname"`
    Age              int       `json:"age"`
    Gender           string    `json:"gender"`
    Username         string    `json:"username"`
    Email            string    `json:"email"`
    PasswordHash     string    `json:"password_hash"`
    RegistrationDate time.Time `json:"registration_date"`
    IsOnline         bool      `json:"is_online"`
}

func GetAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT UserID, Username, is_online FROM User ORDER BY Username")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserID, &user.Username, &user.IsOnline)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
