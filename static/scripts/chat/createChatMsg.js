import { middlePart } from "./config.js";

function createMessageChatBox(username) {
  const messageChatBox = document.createElement("section");
  messageChatBox.classList.add(".messageChatBox");
  const user = document.createElement("h2");
  user.classList.add("receiver-username");
  user.textContent = username;
  const chatDescription = document.createElement("p");
  chatDescription.textContent = `This is the beginning of your direct message history with ${username}`;

  const chatMessages = document.createElement("div");
  chatMessages.classList.add("chatMessages");
  const chatForm = document.createElement("form");

  const chatInput = document.createElement("input");
  chatInput.type = "text";
  chatInput.name = "chatInput";
  chatInput.id = "chatInput";
  chatInput.placeholder = `Message @${username}`;

  const submit = document.createElement("input");
  submit.type = "submit";
  submit.value = "Send";
  chatForm.append(chatInput, submit);

  messageChatBox.append(user, chatDescription, chatMessages, chatForm);

  middlePart.append(messageChatBox);
}

export { createMessageChatBox };
