document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".like-comment-btn, .dislike-comment-btn").forEach(button => {
        button.addEventListener("click", function() {
            const commentID = this.dataset.commentId;
            const isLike = this.classList.contains("like-comment-btn");

            fetch(`/like_dislike_comment?comment_id=${commentID}&is_like=${isLike}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            .then(response => response.json())
            .then(data => {
                document.getElementById(`like-count-${commentID}`).textContent = data.likes;
                document.getElementById(`dislike-count-${commentID}`).textContent = data.dislikes;
            })
            .catch(error => console.error('Error:', error));
        });
    });
});
