import { fetchPrivateMessages } from "./fetchPrivateMessages.js";
import { renderPrivateMessages } from "./renderPrivateMessage.js";

async function displayPrivateMessages() {
  const privateMessagesContainer = document.querySelector(".privateMessages");
  const currentUser = document.querySelector(".username").textContent;
  const selectedUser = document.querySelector(".receiver-username").textContent;

  const messages = await fetchPrivateMessages();
  if (messages.length > 0) {
    renderPrivateMessages(
      messages,
      currentUser,
      selectedUser,
      privateMessagesContainer
    );
  }
}

export { displayPrivateMessages };
