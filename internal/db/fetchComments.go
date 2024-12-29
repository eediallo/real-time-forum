package db

// CreateComments retrieves all comments for a given post ID from the database.
func fetchComments(postID int) ([]Comment, error) {
	commentRows, err := DB.Query(commentsFromDBQuery, postID)
	if err != nil {
		return nil, err
	}
	defer commentRows.Close()

	var comments []Comment
	for commentRows.Next() {
		var comment Comment
		err := commentRows.Scan(&comment.CommentID, &comment.PostID, &comment.Content, &comment.LikeCount, &comment.DislikeCount, &comment.CreatedAt, &comment.Username)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
