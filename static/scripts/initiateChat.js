import { createMessageChatBox } from "./chat/createChatMsg.js";
import { createUserProfile } from "./chat/createUserPro.js";
import { middlePart, leftSide } from "./chat/config.js";
import { setupMessageInputListener } from "./ws.js";

let user = "";
let chatBoxCreated = false;

function initiateChatWithUser(onlineUser) {
  let indexOfDash = onlineUser.textContent.indexOf("-");
  user = onlineUser.textContent.slice(0, indexOfDash);
  middlePart.innerHTML = ""; // clear middle section
  leftSide.innerHTML = "";
  if (!chatBoxCreated) {
    createMessageChatBox(user);
    createUserProfile(user);
    setupMessageInputListener(user);
    chatBoxCreated = false;
  }
}

export { initiateChatWithUser };
