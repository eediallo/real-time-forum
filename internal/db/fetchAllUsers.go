package db

import "time"

func FetchAllUsers() ([]User, error) {
	rows, err := DB.Query(`
        SELECT u.UserID, u.NickName, u.Age, u.FirstName, u.LastName, u.Gender, u.Username, u.Email, u.RegistrationDate, u.is_online, 
        u.LastMessage, COALESCE(MAX(pm.CreatedAt), '1970-01-01 00:00:00') AS LastMessageTime
        FROM User u
        LEFT JOIN PrivateMessages pm ON u.Username = pm.SenderUsername OR u.Username = pm.ReceiverUsername
        GROUP BY u.UserID
        ORDER BY LastMessageTime DESC, u.Username ASC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var lastMessageTimeStr string
		if err := rows.Scan(&user.UserID, &user.NickName, &user.Age, &user.FirstName, &user.LastName, &user.Gender, &user.Username, &user.Email, &user.RegistrationDate, &user.IsOnline, &user.LastMessage, &lastMessageTimeStr); err != nil {
			return nil, err
		}
		user.LastMessageTime, err = time.Parse("2006-01-02 15:04:05", lastMessageTimeStr)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
