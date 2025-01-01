import { initiateChatWithUser } from "../initiateChat.js";
import { onliUsers } from "../chat/config.js";
import { displayPrivateMessages } from "../chat/displayPrivateMessages.js";

function initiateChatHandler() {
  onliUsers.forEach((onlineUser) =>
    onlineUser.addEventListener("click", (e) => {
      initiateChatWithUser(onlineUser);
      displayPrivateMessages();
    })
  );
}
export { initiateChatHandler };
