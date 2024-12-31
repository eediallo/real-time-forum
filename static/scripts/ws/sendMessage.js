import { socket } from "./ws.js";

function sendMessage(username, receiverUsername = null) {
  let jsonData = {
    action: receiverUsername ? "private" : "broadcast",
    message: document.querySelector("#chatInput").value,
    username: username,
    receiverUsername: receiverUsername,
  };
  socket.send(JSON.stringify(jsonData));
  document.querySelector("#chatInput").value = "";
  console.log(jsonData, "<=====json data");
}

export { sendMessage };