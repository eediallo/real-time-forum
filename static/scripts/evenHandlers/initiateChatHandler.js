import { initiateChatWithUser } from "../initiateChat.js";
import { onliUsers } from "../chat/config.js";
import { sendChatContent } from "../chat/sendChatToBackend.js";

function initiateChatHandler() {
  onliUsers.forEach((onlineUser) =>
    onlineUser.addEventListener("click", (e) => {
      initiateChatWithUser(onlineUser);

      const chatForm = document.querySelector("form");
      chatForm.onsubmit = (event) => {
        event.preventDefault();
        sendChatContent();
      };
    })
  );
}

export { initiateChatHandler };
