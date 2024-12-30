import { middlePart } from "./config.js";

function createMessageChatBox(username) {
  const messageChatBox = document.createElement("section");
  messageChatBox.classList.add(".messageChatBox");
  const user = document.createElement("h2");
  user.textContent = username;
  const chatDescription = document.createElement("p");
  chatDescription.textContent = `This is the beginning of your direct message history with ${username}`;

  const chatMessages = document.createElement("div");
  chatMessages.classList.add("chatMessages");
  const chatForm = document.createElement("form");
  chatForm.method = "post";

  const chatInput = document.createElement("input");
  chatInput.type = "text";
  chatInput.name = "chatInput";
  chatInput.id = "chatInput";
  chatInput.placeholder = `Message @${username}`;
  chatForm.append(chatInput);

  messageChatBox.append(user, chatDescription, chatMessages, chatForm);

  middlePart.append(messageChatBox);
}

export { createMessageChatBox };
