package handler

const (
	sessionIDfromCookieQuery = `
			SELECT 
				s.UserID,
				u.Username
			FROM 
				Session AS s
			INNER JOIN
				User AS u
			ON
				s.UserID = u.UserID
			WHERE 
				SessionID = ?`

	userByEmailQuery = `
	SELECT PasswordHash, UserID, NickName, Username
	FROM User
	WHERE Email = ? OR NickName = ?
	`
	deleteSessionByUserIDQuery = `
	DELETE
	FROM 
		Session 
	WHERE 
		UserID = ?
	`
	createNewSessionQuery = `
		INSERT INTO 
			Session
			(
				SessionID, 
				UserID, 
				CreatedAt
			) 
		VALUES (?, ?, ?)`

	addUserDetailsQuery = `
	INSERT INTO 
		User (
			NickName,
			Age,
			FirstName,
			LastName,
			Gender,
			Username, 
			Email, 
			PasswordHash, 
			RegistrationDate
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	insertCommentToDBQuery = `
		INSERT INTO Comments (
			PostID,
			UserID,
			Content,
			CreatedAt
		) VALUES (?, ?, ?, ?)
	`
)

const (
	loginPath     = "/users/login"
	signUpPath    = "/users/sign_up"
	dashboardPath = "/dashboard"
)
