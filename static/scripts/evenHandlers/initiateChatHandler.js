import { initiateChatWithUser } from "../initiateChat.js";
import { onliUsers } from "../chat/config.js";

function initiateChatHandler() {
  onliUsers.forEach((onlineUser) =>
    onlineUser.addEventListener("click", (e) => {
      initiateChatWithUser(onlineUser);
    })
  );
}

export { initiateChatHandler };
