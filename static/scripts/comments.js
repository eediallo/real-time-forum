document.addEventListener("DOMContentLoaded", function () {
  document.querySelectorAll(".comment-btn").forEach(function (button) {
    button.addEventListener("click", function () {
      var postID = this.getAttribute("data-post-id");
      var commentsSection = document.getElementById(`comments-${postID}`);
      commentsSection.style.display =
        commentsSection.style.display === "none" ? "block" : "none";
    });
  });
});

function submitComment(event, postID) {
  event.preventDefault();
  var form = event.target;
  var formData = new FormData(form);
  fetch(form.action, {
    method: form.method,
    body: formData,
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.success) {
        var commentsSection = document.getElementById(`comments-${postID}`);
        var newComment = document.createElement("div");
        newComment.classList.add("comment");
        newComment.innerHTML = `<p>${data.comment.Content}</p>
                                            <p><small>Commented by: ${data.comment.Username} on: ${data.comment.CreatedAt}</small></p>`;
        commentsSection.insertBefore(newComment, form);
        form.reset();
        document.getElementById(`comment-count-${postID}`).textContent =
          data.commentCount;
      }
    });
}
