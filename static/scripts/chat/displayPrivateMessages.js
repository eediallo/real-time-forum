import { fetchPrivateMessages } from "./fetchPrivateMessages.js";

async function displayPrivateMessages() {
  const privateMessagesContainer = document.querySelector(".privateMessages");
  const currentUser = document.querySelector(".username").textContent;
  const selectedUser = document.querySelector(".receiver-username").textContent;

  const messages = await fetchPrivateMessages();
  if (messages.length > 0) {
    renderPrivateMessages();
  }

  function renderPrivateMessages() {
    const filteredMessages = messages.filter(
      (message) =>
        (message.senderUsername === currentUser &&
          message.receiverUsername === selectedUser) ||
        (message.senderUsername === selectedUser &&
          message.receiverUsername === currentUser)
    );

    const groupedMessages = groupMessagesByDate(filteredMessages);

    for (const [date, messages] of Object.entries(groupedMessages)) {
      const dateElement = document.createElement("div");
      dateElement.classList.add("date");
      dateElement.innerHTML = `<p><strong>${date}</strong></p>`;
      privateMessagesContainer.appendChild(dateElement);

      messages.forEach((message) => {
        const messageElement = document.createElement("div");
        messageElement.classList.add("message");
        messageElement.innerHTML = `
                                                                    <p><strong>${message.senderUsername}</strong> <small>${message.createdAt}</small></p>
                                                                    <p>${message.content}</p>
                                                                    
                                                            `;
        privateMessagesContainer.appendChild(messageElement);
      });
    }
  }

  function groupMessagesByDate(messages) {
    return messages.reduce((acc, message) => {
      const date = new Date(message.createdAt).toLocaleDateString();
      if (!acc[date]) {
        acc[date] = [];
      }
      acc[date].push(message);
      return acc;
    }, {});
  }
}

export { displayPrivateMessages };
