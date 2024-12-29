package db

import (
	"fmt"
)

// CreatePosts retrieves all posts from the database along with their associated comments.
func FetchPosts() ([]Post, error) {
	postRows, err := DB.Query(postsWithDetailsQuery)
	if err != nil {
		return nil, err
	}
	defer postRows.Close()

	var posts []Post
	for postRows.Next() {
		var post Post
		err := postRows.Scan(&post.PostID, &post.Title, &post.Content, &post.CommentCount, &post.CreatedAt, &post.Category, &post.Username)
		if err != nil {
			return nil, err
		}

		// Fetch like and dislike counts for the post
		post.LikeCount, post.DislikeCount, _ = GetLikeDislikeCounts(post.PostID)

		comments, err := fetchComments(post.PostID)
		if err != nil {
			return nil, fmt.Errorf("error copying colums into row ---CreatePosts--%s", err.Error())
		}

		// Fetch comment count for the post
		err = DB.QueryRow("SELECT COUNT(*) FROM comments WHERE PostID = ?", post.PostID).Scan(&post.CommentCount)
		if err != nil {
			return nil, err
		}

		post.Comments = comments
		posts = append(posts, post)
	}

	return posts, nil
}
