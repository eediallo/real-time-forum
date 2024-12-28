let socket = null;
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
    console.log(msg);
    let j = JSON.parse(msg.data);
    console.log(j);
  };
});

const onlineUser = document.querySelector(".online");
console.log(onlineUser);

onlineUser.addEventListener("click", () => {
  console.log("User name clicked...........");
  let jsonData = {};
  jsonData["action"] = "username";
  jsonData["username"] = onlineUser.textContent;
  socket.send(JSON.stringify(jsonData));
});
