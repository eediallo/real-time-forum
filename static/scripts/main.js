import { eventHandlers } from "./eventHandlers.js";

function main() {
  eventHandlers();
}

//================================================================

const onliUsers = document.querySelectorAll(".online");
const dashboardContainer = document.querySelector(".dashboard-container");
const middlePart = document.querySelector(".middle-part");
let user = "";

onliUsers.forEach((onlineUser) =>
  onlineUser.addEventListener("click", (e) => {
    let indexOfDash = onlineUser.textContent.indexOf("-");
    user += onlineUser.textContent.slice(0, indexOfDash);
    middlePart.innerHTML = ""; // clear dashboard container
    createMessageChatBox(user);
  })
);

function createMessageChatBox(username) {
  const messageChatBox = document.createElement("section");
  const user = document.createElement("h2");
  user.textContent = username;
  const chatDescription = document.createElement("p");
  chatDescription.textContent = `This is the beginning of your direct message history with ${username}`;

  const chatInput = document.createElement("input");
  chatInput.type = "text";
  chatInput.name = "chatInput";
  chatInput.id = "chatInput";

  messageChatBox.append(user, chatDescription, chatInput);
  middlePart.append(messageChatBox);
}

window.onload = main;
