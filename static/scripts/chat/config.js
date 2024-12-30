const onliUsers = document.querySelectorAll(".online");
const dashboardContainer = document.querySelector(".dashboard-container");
const middlePart = document.querySelector(".middle-part");
const leftSide = document.querySelector(".left-side");
const userTextingMsg = document.querySelector(".username").textContent;
const password = document.querySelector("#password").value;
const signUpBtn = document.querySelector("#signUpBtn");

//layout change
const postContainer = document.querySelector(".post-container");
postContainer.style.display = "none";

// event to display post creator
const whatOnYourMind = document.querySelector("#whatOnYourMind");
whatOnYourMind.addEventListener("click", () => {
  postContainer.style.display = "block";
});

export {
  onliUsers,
  dashboardContainer,
  middlePart,
  leftSide,
  userTextingMsg,
  password,
  signUpBtn,
};
