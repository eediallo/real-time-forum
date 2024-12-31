import { groupMessagesByDate } from "./groupeMsgByDate.js";

function renderPrivateMessages(
  messages,
  currentUser,
  selectedUser,
  privateMessagesContainer
) {
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

export { renderPrivateMessages };
