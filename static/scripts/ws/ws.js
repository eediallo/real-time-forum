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
        const chatMessages = document.querySelector(".PrivateMessages");
        chatMessages.innerHTML += `${data.message}<br>`;
        break;
      // case "private":
      //   const privateMessages = document.querySelector(".privateMessages");
      //   privateMessages.innerHTML += `${data.message}<br>`;
      //   break;
    }
  };
});

export { socket };
