import { createMessageChatBox } from "./chat/createChatMsg.js";
import { createUserProfile } from "./chat/createUserPro.js";
import { middlePart, leftSide } from "./chat/config.js";
import { setupMessageInputListener } from "./ws/inputInitializer.js";

let userReceivingMsg = "";
const userTextingMsg = document.querySelector(".username").textContent;
let chatBoxCreated = false;

function initiateChatWithUser(onlineUser) {
  let indexOfDash = onlineUser.textContent.indexOf("-");
  userReceivingMsg = onlineUser.textContent.slice(0, indexOfDash);
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
