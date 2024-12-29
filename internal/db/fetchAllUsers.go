package db

func FetchAllUsers() ([]User, error) {
	rows, err := DB.Query("SELECT UserID, NickName, Age, FirstName, LastName, Gender, Username, Email, RegistrationDate, is_online FROM User WHERE is_online = 1 OR is_online = 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.NickName, &user.Age, &user.FirstName, &user.LastName, &user.Gender, &user.Username, &user.Email, &user.RegistrationDate, &user.IsOnline); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
