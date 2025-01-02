import { sendChatContent } from "../chat/sendChatToBackend.js";
import { sendMessage, showNotification } from "./sendMessage.js";
import { socket } from "./ws.js";

function setupMessageInputListener(username) {
  const messageInput = document.querySelector("#chatInput");
  const receiverUsername =
    document.querySelector(".receiver-username").textContent;

  messageInput.addEventListener("keydown", (e) => {
    if (e.code === "Enter") {
      if (!socket) {
        console.log("No connection");
        return false;
      }
      e.preventDefault();
      const chatInputValue = messageInput.value;
      sendMessage(username, receiverUsername);
      showNotification();
      sendChatContent(chatInputValue);
    }
  });

  const chatForm = document.querySelector("form");
  chatForm.onsubmit = (event) => {
    event.preventDefault();
    const chatInputValue = messageInput.value;
    sendMessage(username, receiverUsername);
    sendChatContent(chatInputValue);
  };
}

export { setupMessageInputListener };
