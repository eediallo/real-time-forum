const chatMessages = document.querySelector(".chat-messages");

let socket = null;
window.onbeforeunload = () => {
  console.log("leaving.......");
  let jsonData = {};
  jsonData["action"] = "left";
  socket.send(JSON.stringify(jsonData));
};

document.addEventListener("DOMContentLoaded", () => {
  socket = new WebSocket("ws://127.0.0:8080/ws");

  socket.onopen = () => {
    console.log("Successully connected");
  };

  socket.onclose = () => {
    console.log("Connection closed");
  };

  socket.onerror = (error) => {
    console.log("there was an error");
  };

  socket.onmessage = (msg) => {
    let data = JSON.parse(msg.data);
    let jsonData = {};
    jsonData["action"] = "username";
    socket.send(JSON.stringify(jsonData));
    switch (data.action) {
      case "broadcast":
        chatMessages.innerHTML = chatMessages.innerHTML + data.message + "<br>";
        break;
    }
  };
});

const onlineUsers = document.querySelectorAll(".online-user");

onlineUsers.forEach((onlineUser) => {
  onlineUser.addEventListener("click", () => {
    createChatBox();
    setupMessageInputListener();
  });
});

function setupMessageInputListener() {
  const messageTextArea = document.querySelector(".message");
  messageTextArea.addEventListener("keydown", (e) => {
    if (e.code === "Enter") {
      if (!socket) {
        console.log("no connection");
        return false;
      }
      e.preventDefault();
      e.stopPropagation();
      sendMessage();
    }
  });
}

function createChatBox() {
  if (!document.querySelector(".chat-box")) {
    const chatBoxCard = document.createElement("section");
    const chatBox = document.createElement("div");
    chatBox.classList.add("chat-box");
    const chatBoxInput = document.createElement("input");
    chatBoxInput.classList.add("message");
    chatBoxInput.type = "textArea";
    chatBoxInput.placeholder = "Enter your message";
    const sendBtn = document.createElement("button");
    sendBtn.classList.add("sentMsgBtn");
    sendBtn.textContent = "Send";
    chatBox.append(chatBoxInput, sendBtn);
    chatBoxCard.append(chatBox);

    document.body.appendChild(chatBoxCard);
  }
}

function sendMessage() {
  let jsonData = {};
  jsonData["action"] = "broadcast";
  onlineUsers.forEach((onlineUser) => {
    jsonData["username"] = onlineUser.textContent;
  });
  jsonData["message"] = document.querySelector(".message").value;
  socket.send(JSON.stringify(jsonData));
  document.querySelector(".message").value = "";
  console.log(jsonData, "<=====json data");
}
