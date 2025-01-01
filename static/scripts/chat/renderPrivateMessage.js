import { groupMessagesByDate } from "./groupeMsgByDate.js";

let loadedMessagesCount = 0;
const messagesPerLoad = 10;

function renderPrivateMessages(
  messages,
  currentUser,
  selectedUser,
  privateMessagesContainer,
  loadMore = false,
  loadDirection = "down"
) {
  if (!loadMore) {
    privateMessagesContainer.innerHTML = "";
    loadedMessagesCount = 0;
  }

  const filteredMessages = messages.filter(
    (message) =>
      (message.senderUsername === currentUser &&
        message.receiverUsername === selectedUser) ||
      (message.senderUsername === selectedUser &&
        message.receiverUsername === currentUser)
  );

  let messagesToRender;
  if (loadDirection === "up") {
    messagesToRender = filteredMessages.slice(
      -loadedMessagesCount - messagesPerLoad,
      -loadedMessagesCount || undefined
    );
  } else {
    messagesToRender = filteredMessages.slice(
      loadedMessagesCount,
      loadedMessagesCount + messagesPerLoad
    );
  }

  const groupedMessages = groupMessagesByDate(messagesToRender);
  loadedMessagesCount += messagesToRender.length;

  const initialScrollHeight = privateMessagesContainer.scrollHeight;

  for (const [date, messages] of Object.entries(groupedMessages)) {
    const dateElement = document.createElement("div");
    dateElement.classList.add("date");
    dateElement.innerHTML = `<p><strong>${date}</strong></p>`;
    if (loadDirection === "up") {
      privateMessagesContainer.prepend(dateElement);
    } else {
      privateMessagesContainer.append(dateElement);
    }

    messages.forEach((message) => {
      const messageElement = document.createElement("div");
      messageElement.classList.add("message");
      messageElement.innerHTML = `
                <p><strong>${message.senderUsername}</strong> <small>${message.createdAt}</small></p>
                <p>${message.content}</p>
            `;
      if (loadDirection === "up") {
        privateMessagesContainer.prepend(messageElement);
      } else {
        privateMessagesContainer.append(messageElement);
      }
    });
  }

  if (!loadMore) {
    privateMessagesContainer.scrollTop = privateMessagesContainer.scrollHeight;
  } else if (loadMore && loadDirection === "up") {
    privateMessagesContainer.scrollTop =
      privateMessagesContainer.scrollHeight - initialScrollHeight;
  }
}

function debounce(func, wait) {
  let timeout;
  return function (...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply(this, args), wait);
  };
}

function setupScrollListener(
  messages,
  currentUser,
  selectedUser,
  privateMessagesContainer
) {
  const handleScroll = debounce(() => {
    if (privateMessagesContainer.scrollTop === 0) {
      renderPrivateMessages(
        messages,
        currentUser,
        selectedUser,
        privateMessagesContainer,
        true,
        "up"
      );
    } else if (
      privateMessagesContainer.scrollTop +
        privateMessagesContainer.clientHeight >=
      privateMessagesContainer.scrollHeight
    ) {
      renderPrivateMessages(
        messages,
        currentUser,
        selectedUser,
        privateMessagesContainer,
        true,
        "down"
      );
    }
  }, 200);

  privateMessagesContainer.addEventListener("scroll", handleScroll);
}

export { renderPrivateMessages, setupScrollListener };
