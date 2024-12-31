import { initiateChatWithUser } from "../initiateChat.js";
import { onliUsers } from "../chat/config.js";
import { sendChatContent } from "../chat/sendChatToBackend.js";

function initiateChatHandler() {
  onliUsers.forEach((onlineUser) =>
    onlineUser.addEventListener("click", (e) => {
      initiateChatWithUser(onlineUser);

      //display private messages ===========

      const privateMessagesContainer =
        document.querySelector(".privateMessages");

      if (privateMessagesContainer) {
        const currentUser = document.querySelector(".username").textContent;
        console.log(currentUser);
        const selectedUser =
          document.querySelector(".receiver-username").textContent;
        console.log(selectedUser);

        fetch("/private_messages")
          .then((response) => response.json())
          .then((messages) => {
            privateMessagesContainer.innerHTML = ""; // Clear previous messages
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
                  <p>${message.content}</p>
                  <p><small>${message.createdAt}</small></p>
                `;
                privateMessagesContainer.appendChild(messageElement);
              });
          })
          .catch((error) => {
            console.error("Error fetching private messages:", error);
          });
      }
    })
  );
}

export { initiateChatHandler };
