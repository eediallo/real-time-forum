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
    // console.log(msg);
    // let j = JSON.parse(msg.data);
    // console.log(j);

    let data = JSON.parse(msg.data);
    console.log("Action is ", data.action);
    let jsonData = {};
    jsonData["action"] = "username";
    socket.send(JSON.stringify(jsonData));

    switch (data.action) {
      case "list_users":
        break;
    }
  };
});

const onlineUsers = document.querySelectorAll(".online-user");

onlineUsers.forEach((onlineUser) => {
  console.log(onlineUser, "<---online user");
  onlineUser.addEventListener("click", () => {
    console.log(`User name: ${onlineUser.textContent} clicked...........`);
    chatBox();
  });
});

function chatBox() {
  if (!document.querySelector(".chat-box")) {
    const chatBoxCard = document.createElement("section");
    const chatBox = document.createElement("div");
    chatBox.classList.add("chat-box");
    const chatBoxInput = document.createElement("input");
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

// <!-- Private Messages Section -->
// <div class="private-messages">
//   <h3>Private Messages</h3>
//   <div class="user-list">
//     <h4>Users</h4>
//     <ul id="user-list">
//       {{range .OnlineUsers}}
//       <li
//         data-username="{{.Username}}"
//         class="{{if .IsOnline}}online{{else}}offline{{end}}"
//       >
//         {{.Username}}
//         <span class="last-message">{{.LastMessage}}</span>
//       </li>
//       {{else}}
//       <li>No users available</li>
//       {{end}}
//     </ul>
//   </div>

{
  /* <div class="chat-box">
        <div class="chat-header">
          <h4 id="chat-with">Chat with: <span id="chat-username"></span></h4>
        </div>
        <div class="chat-messages" id="private-chat-messages"></div>
        <div class="chat-input">
          <input
            type="text"
            id="private-chat-input"
            placeholder="Type a message..."
          />
          <button id="private-chat-send">Send</button>
        </div>
      </div>


      <div id="chat-box-output"></div> */
}
