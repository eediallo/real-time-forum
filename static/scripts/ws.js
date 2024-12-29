let socket = null;

window.onbeforeunload = () => {
  if (socket) {
    console.log("leaving.......");
    let jsonData = { action: "left" };
    socket.send(JSON.stringify(jsonData));
  }
};

document.addEventListener("DOMContentLoaded", () => {
  socket = new WebSocket("ws://127.0.0.1:8080/ws");

  socket.onopen = () => {
    console.log("Successfully connected");
  };

  socket.onclose = () => {
    console.log("Connection closed");
  };

  socket.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  socket.onmessage = (msg) => {
    let data;
    try {
      data = JSON.parse(msg.data);
    } catch (e) {
      console.error("Error parsing JSON:", e);
      return;
    }
    switch (data.action) {
      case "broadcast":
        const chatMessages = document.querySelector(".chatMessages");
        chatMessages.innerHTML += `${data.message}<br>`;
        break;
    }
  };
});

// const onlineUsers = document.querySelectorAll(".online-user");

// onlineUsers.forEach((onlineUser) => {
//   onlineUser.addEventListener("click", () => {
//     createChatBox();
//     setupMessageInputListener(onlineUser.textContent);
//   });
// });

function setupMessageInputListener(username) {
  const messageTextArea = document.querySelector("#chatInput");
  messageTextArea.addEventListener("keydown", (e) => {
    if (e.code === "Enter") {
      if (!socket) {
        console.log("No connection");
        return false;
      }
      e.preventDefault();
      sendMessage(username);
    }
  });
}

// function createChatBox() {
//   if (!document.querySelector(".chat-box")) {
//     const chatBoxCard = document.createElement("section");
//     const chatBox = document.createElement("div");
//     chatBox.classList.add("chat-box");
//     const chatBoxInput = document.createElement("input");
//     chatBoxInput.classList.add("message");
//     chatBoxInput.type = "text";
//     chatBoxInput.placeholder = "Enter your message";
//     const sendBtn = document.createElement("button");
//     sendBtn.classList.add("sentMsgBtn");
//     sendBtn.textContent = "Send";
//     chatBox.append(chatBoxInput, sendBtn);
//     chatBoxCard.append(chatBox);

//     document.body.appendChild(chatBoxCard);
//   }
// }

function sendMessage(username) {
  let jsonData = {
    action: "broadcast",
    message: document.querySelector("#chatInput").value,
    username: username,
  };
  socket.send(JSON.stringify(jsonData));
  document.querySelector("#chatInput").value = "";
  console.log(jsonData, "<=====json data");
}

//layout change
const postContainer = document.querySelector(".post-container");
postContainer.style.display = "none";

// event to display post creator
const whatOnYourMind = document.querySelector("#whatOnYourMind");
whatOnYourMind.addEventListener("click", () => {
  postContainer.style.display = "block";
});

export { setupMessageInputListener };
