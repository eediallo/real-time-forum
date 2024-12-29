import { leftSide } from "./config.js";

function createUserProfile(username) {
  const profileSection = document.createElement("section");
  const user = document.createElement("h2");
  user.textContent = username;
  const userProfile = `http://localhost:8080/users/${username}`;
  const a = document.createElement("a");
  a.textContent = `${username}#`;
  a.setAttribute("href", userProfile);

  profileSection.append(user, a);
  leftSide.append(profileSection);
}

export { createUserProfile };
