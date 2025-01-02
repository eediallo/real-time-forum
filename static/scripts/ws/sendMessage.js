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

async function showNotification() {
  const permission = await Notification.requestPermission();

  if (permission === "granted") {
    const Notification = new Notification("New message", {
      body: "You have recieved a new message",
    });
  }
}

export { sendMessage, showNotification };
