package db

type Post struct {
	UserID       int
	PostID       int
	Title        string
	Content      string
	CreatedAt    string
	Category     string
	Username     string
	MediaPath    string
	Comments     []Comment
	CommentCount int
	LikeCount    int
	DislikeCount int
}
