package db

var (
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
)
