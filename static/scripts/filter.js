function filterPosts(category) {
  const posts = document.querySelectorAll(".post"); // get all posts
  posts.forEach((post) => {
    const postCategory = post.querySelector(".category").textContent; // get the category of the current post

    // Check if the post's category matches the selected category or if "All" is selected
    if (category === "All" || postCategory === category) {
      post.style.display = "block"; // Show the post
    } else {
      post.style.display = "none"; // Hide the post
    }
  });
}

// Add event listeners to the filter buttons
const filterButtons = document.querySelectorAll(".filter-container .btn");
filterButtons.forEach((button) => {
  button.addEventListener("click", function (event) {
    // Remove 'active' class from all buttons
    filterButtons.forEach((btn) => btn.classList.remove("active"));

    // Add 'active' class to the clicked button
    this.classList.add("active");

    // Get the category from the data-category attribute of the clicked button
    const category = this.getAttribute("data-category");

    // Call the function to filter posts based on the selected category
    filterPosts(category);
  });
});

// Initialize to show all posts by default
document.addEventListener("DOMContentLoaded", () => {
  filterPosts("All");
});
