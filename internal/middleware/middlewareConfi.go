package middleware

const (
	userByUsernameAndUserIDQuery = `
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
)

const (
	signUpPath = "/users/sign_up"
	loginPath  = "/users/login"
)
