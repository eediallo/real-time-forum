import { createMessageChatBox } from "./chat/createChatMsg.js";
import { createUserProfile } from "./chat/createUserPro.js";
import { middlePart, leftSide } from "./chat/config.js";

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
    chatBoxCreated = false;
  }
}

export { initiateChatWithUser };
