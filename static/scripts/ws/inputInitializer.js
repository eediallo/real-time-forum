import { sendMessage } from "./sendMessage.js";
import { socket } from "./ws.js";
socket;
function setupMessageInputListener(username) {
  const messageInput = document.querySelector("#chatInput");
  messageInput.addEventListener("keydown", (e) => {
    if (e.code === "Enter") {
      if (!socket) {
        console.log("No connection");
        return false;
      }
      e.preventDefault();
      sendMessage(username);
    }
  });
}

export { setupMessageInputListener };
