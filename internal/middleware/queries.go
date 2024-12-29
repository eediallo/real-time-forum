package middleware

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
)
