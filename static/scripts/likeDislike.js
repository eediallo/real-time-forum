document.addEventListener("DOMContentLoaded", function() {
    document.querySelectorAll(".like-btn").forEach(function(button) {
        button.addEventListener("click", function() {
            var postID = this.getAttribute("data-post-id");
            fetch(`/like?post_id=${postID}`, {
                method: "POST"
            }).then(response => response.json())
            .then(data => {
                document.getElementById(`like-count-${postID}`).textContent = data.likes;
                document.getElementById(`dislike-count-${postID}`).textContent = data.dislikes;
            });
        });
    });

    document.querySelectorAll(".dislike-btn").forEach(function(button) {
        button.addEventListener("click", function() {
            var postID = this.getAttribute("data-post-id");
            fetch(`/dislike?post_id=${postID}`, {
                method: "POST"
            }).then(response => response.json())
            .then(data => {
                document.getElementById(`like-count-${postID}`).textContent = data.likes;
                document.getElementById(`dislike-count-${postID}`).textContent = data.dislikes;
            });
        });
    });
});