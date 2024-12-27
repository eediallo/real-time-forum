document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.share-btn').forEach(function(button) {
        button.addEventListener('click', function() {
            var postID = this.getAttribute('data-post-id');
            // Add functionality to share the post here
            alert('Shared post ' + postID);
        });
    });
})