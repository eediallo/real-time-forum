package db

func GetUserByUsername(username string) (User, error) {
	query := `
        SELECT 
            UserID, NickName, Age, FirstName, LastName, Gender, Username, Email, RegistrationDate, is_online
        FROM 
            User 
        WHERE 
            Username = ?`

	row := DB.QueryRow(query, username)

	var user User
	if err := row.Scan(&user.UserID, &user.NickName, &user.Age, &user.FirstName, &user.LastName, &user.Gender, &user.Username, &user.Email, &user.RegistrationDate, &user.IsOnline); err != nil {
		return User{}, err
	}

	return user, nil
}
