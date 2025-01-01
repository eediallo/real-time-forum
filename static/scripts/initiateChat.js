import { createMessageChatBox } from "./chat/createChatMsg.js";
import { createUserProfile } from "./chat/createUserPro.js";
import { middlePart, leftSide, userTextingMsg } from "./chat/config.js";
import { setupMessageInputListener } from "./ws/inputInitializer.js";

let userReceivingMsg = "";
let chatBoxCreated = false;

function initiateChatWithUser(onlineUser) {
  let indexOfDash = onlineUser.textContent.indexOf("-");
  userReceivingMsg = onlineUser.textContent.slice(0, indexOfDash);
  userReceivingMsg = userReceivingMsg.trim();
  middlePart.innerHTML = ""; // clear middle section
  leftSide.innerHTML = "";
  if (!chatBoxCreated) {
    createMessageChatBox(userReceivingMsg);
    createUserProfile(userReceivingMsg);
    setupMessageInputListener(userTextingMsg);
    chatBoxCreated = false;
  }
}

export { initiateChatWithUser };
