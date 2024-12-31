import { socket } from "./ws.js";

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

export { sendMessage };
