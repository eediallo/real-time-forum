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

	postsWithDetailsQuery = `
	SELECT
		p.PostID,
		p.Title, 
		p.Content,
		p.CommentCount,
		STRFTIME('%d/%m/%Y, %H:%M', p.CreatedAt) AS CreatedAt,
		p.Category,
		u.Username
	FROM 
		Post AS p
	INNER JOIN
		User AS u
	ON
		p.UserID = u.UserID
	ORDER BY
		p.CreatedAt DESC
	`
	commentsFromDBQuery = `
	SELECT
		c.CommentID,
		c.PostID,
		c.Content,
		c.LikeCount,
		c.DislikeCount,
		STRFTIME('%d/%m/%Y, %H:%M', c.CreatedAt) AS CreatedAt,
		u.Username
	FROM
		Comments AS c
	INNER JOIN
		User AS u
	ON
		c.UserID = u.UserID
	WHERE
		c.PostID = ?
	ORDER BY
		(c.CreatedAt) ASC
	`
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
