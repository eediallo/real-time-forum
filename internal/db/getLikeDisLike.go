package db

func GetLikeDislikeCounts(postID int) (int, int, error) {
	var likeCount, dislikeCount int
	err := DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE PostID = ? AND IsLike = 1", postID).Scan(&likeCount)
	if err != nil {
		return 0, 0, err
	}
	err = DB.QueryRow("SELECT COUNT(*) FROM LikeDislike WHERE PostID = ? AND IsLike = 0", postID).Scan(&dislikeCount)
	if err != nil {
		return 0, 0, err
	}
	return likeCount, dislikeCount, nil
}
