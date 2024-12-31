import { sendMessage } from "./sendMessage.js";
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
      sendMessage(username, receiverUsername);
    }
  });
}

export { setupMessageInputListener };
