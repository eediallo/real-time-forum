package db

type Comment struct {
	CommentID    int
	PostID       int
	Content      string
	CreatedAt    string
	Username     string
	LikeCount    int
	DislikeCount int
}
