import { leftSide } from "./config.js";

function createUserProfile(username) {
  const profileSection = document.createElement("section");
  const user = document.createElement("h2");
  user.textContent = username;
  const userProfile = `http://localhost:8080/users/${username}`;
  const profileTag = document.createElement("p");
  profileTag.textContent = `${username}#`;

  const fullProfileLink = document.createElement("a");
  fullProfileLink.setAttribute("href", userProfile);
  fullProfileLink.textContent = "fullProfile#";

  profileSection.append(user, profileTag, fullProfileLink);
  leftSide.append(profileSection);
}

export { createUserProfile };
