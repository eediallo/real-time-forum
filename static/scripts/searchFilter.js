const state = {
  searchTerm: "",
  filteredTerms: [],
};

function searchInputHandler(e) {
  state.searchTerm = e.target.value.toLowerCase();
  const posts = document.querySelectorAll(".post");

  posts.forEach((post) => {
    const title = post.querySelector("h3").textContent.toLowerCase();
    const descriptions = post.querySelectorAll("p");
    let descriptionText = "";
    descriptions.forEach((desc) => {
      descriptionText += desc.textContent.toLowerCase() + " ";
    });

    if (
      title.includes(state.searchTerm) ||
      descriptionText.includes(state.searchTerm)
    ) {
      post.style.display = "block";
    } else {
      post.style.display = "none";
    }
  });
}

export { searchInputHandler };
