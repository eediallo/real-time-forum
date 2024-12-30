// // JavaScript for private messaging functionality
// document.querySelectorAll("#user-list li").forEach((user) => {
//   user.addEventListener("click", function () {
//     const username = this.dataset.username;
//     document.getElementById("chat-username").textContent = username;
//     loadMessages(username);
//   });
// });

// document
//   .getElementById("private-chat-send")
//   .addEventListener("click", function () {
//     const input = document.getElementById("private-chat-input");
//     const message = input.value;
//     const username = document.getElementById("chat-username").textContent;
//     if (message.trim() !== "") {
//       sendMessage(username, message);
//       input.value = "";
//     }
//   });

// function loadMessages(username) {
//   // Fetch and display the last 10 messages with the user
//   // Implement throttling for loading more messages on scroll
//   const chatMessages = document.getElementById("private-chat-messages");
//   chatMessages.innerHTML = ""; // Clear previous messages
//   fetch(`/load_messages?username=${username}&limit=10`)
//     .then((response) => response.json())
//     .then((messages) => {
//       messages.forEach((message) => {
//         const messageElement = document.createElement("div");
//         messageElement.classList.add("message");
//         messageElement.innerHTML = `<p><strong>${message.username}</strong> [${message.date}]: ${message.content}</p>`;
//         chatMessages.appendChild(messageElement);
//       });
//     });

//   chatMessages.addEventListener(
//     "scroll",
//     throttle(function () {
//       if (chatMessages.scrollTop === 0) {
//         loadMoreMessages(username);
//       }
//     }, 200)
//   );
// }

// function loadMoreMessages(username) {
//   // Fetch and display 10 more messages with the user
//   const chatMessages = document.getElementById("private-chat-messages");
//   const currentMessageCount = chatMessages.children.length;
//   fetch(
//     `/load_messages?username=${username}&limit=10&offset=${currentMessageCount}`
//   )
//     .then((response) => response.json())
//     .then((messages) => {
//       messages.forEach((message) => {
//         const messageElement = document.createElement("div");
//         messageElement.classList.add("message");
//         messageElement.innerHTML = `<p><strong>${message.username}</strong> [${message.date}]: ${message.content}</p>`;
//         chatMessages.insertBefore(messageElement, chatMessages.firstChild);
//       });
//     });
// }

// function sendMessage(username, message) {
//   // Send the message to the server and update the chat
//   fetch(`/send_message`, {
//     method: "POST",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify({ username, message }),
//   })
//     .then((response) => response.json())
//     .then((data) => {
//       if (data.success) {
//         const chatMessages = document.getElementById("private-chat-messages");
//         const messageElement = document.createElement("div");
//         messageElement.classList.add("message");
//         messageElement.innerHTML = `<p><strong>${data.message.username}</strong> [${data.message.date}]: ${data.message.content}</p>`;
//         chatMessages.appendChild(messageElement);
//       }
//     });
// }

// function throttle(func, limit) {
//   let lastFunc;
//   let lastRan;
//   return function () {
//     const context = this;
//     const args = arguments;
//     if (!lastRan) {
//       func.apply(context, args);
//       lastRan = Date.now();
//     } else {
//       clearTimeout(lastFunc);
//       lastFunc = setTimeout(function () {
//         if (Date.now() - lastRan >= limit) {
//           func.apply(context, args);
//           lastRan = Date.now();
//         }
//       }, limit - (Date.now() - lastRan));
//     }
//   };
// }

// <!-- Online Users Section -->
// <div class="online-users">
//   <h3>Online Users</h3>
//   <ul>
//     {{range .OnlineUsers}}
//     <li class="online-user">{{.Username}}</li>
//     {{else}}
//     <li>No users online</li>
//     {{end}}
// //   </ul>
// </div>




// <div class="chat-messages"></div>
