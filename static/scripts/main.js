import { eventHandlers } from "./eventHandlers.js";

function main() {
  eventHandlers();
}

//================================================================

const onliUsers = document.querySelectorAll(".online");
const dashboardContainer = document.querySelector(".dashboard-container");
const middlePart = document.querySelector(".middle-part");
const leftSide = document.querySelector(".left-side");
let user = "";

let chatBoxCreated = false;

onliUsers.forEach((onlineUser) =>
  onlineUser.addEventListener("click", (e) => {
    let indexOfDash = onlineUser.textContent.indexOf("-");
    user = onlineUser.textContent.slice(0, indexOfDash);
    middlePart.innerHTML = ""; // clear middle section
    leftSide.innerHTML = "";
    if (!chatBoxCreated) {
      createMessageChatBox(user);
      createUserProfile(user);
      chatBoxCreated = false;
    }
  })
);

function createMessageChatBox(username) {
  const messageChatBox = document.createElement("section");
  messageChatBox.classList.add(".messageChatBox");
  const user = document.createElement("h2");
  user.textContent = username;
  const chatDescription = document.createElement("p");
  chatDescription.textContent = `This is the beginning of your direct message history with ${username}`;

  const chatInput = document.createElement("input");
  chatInput.type = "text";
  chatInput.name = "chatInput";
  chatInput.id = "chatInput";
  chatInput.placeholder = `Message @${username}`;

  messageChatBox.append(user, chatDescription, chatInput);

  middlePart.append(messageChatBox);
}

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

window.onload = main;
