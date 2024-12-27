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
	LastMessage      string    `json:"last_message"`
}

// func GetAllUsers() ([]User, error) {
// 	rows, err := DB.Query(`
// 		SELECT u.UserID, u.Username, u.is_online, COALESCE(pm.Content, '') AS LastMessage
// 		FROM User u
// 		LEFT JOIN (
// 			SELECT p1.*
// 			FROM PrivateMessages p1
// 			INNER JOIN (
// 				SELECT MAX(MessageID) AS MessageID
// 				FROM PrivateMessages
// 				GROUP BY SenderID, ReceiverID
// 			) p2 ON p1.MessageID = p2.MessageID
// 		) pm ON u.UserID = pm.SenderID OR u.UserID = pm.ReceiverID
// 		ORDER BY u.is_online DESC, pm.CreatedAt DESC, u.Username
// 	`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.UserID, &user.Username, &user.IsOnline, &user.LastMessage)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}
// 	return users, nil
// }
