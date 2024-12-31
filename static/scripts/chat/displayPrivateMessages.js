import { fetchPrivateMessages } from "./fetchPrivateMessages.js";

async function displayPrivateMessages() {
  const privateMessagesContainer = document.querySelector(".privateMessages");
  const currentUser = document.querySelector(".username").textContent;
  const selectedUser = document.querySelector(".receiver-username").textContent;
  const messages = await fetchPrivateMessages();
  renderPrivateMessages();

  function renderPrivateMessages() {
    messages
      .filter(
        (message) =>
          (message.senderUsername === currentUser &&
            message.receiverUsername === selectedUser) ||
          (message.senderUsername === selectedUser &&
            message.receiverUsername === currentUser)
      )
      .forEach((message) => {
        const messageElement = document.createElement("div");
        messageElement.classList.add("message");
        messageElement.innerHTML = `
                <p><strong>${message.senderUsername}</strong>: ${message.content}</p>
                <p><small>${message.createdAt}</small></p>
              `;
        privateMessagesContainer.appendChild(messageElement);
      });
  }
}

export { displayPrivateMessages };
