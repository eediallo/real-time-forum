let socket = null;
document.addEventListener("DOMContentLoaded", () => {
  socket = new WebSocket("ws://127.0.0:8080/ws");

  socket.onopen = () => {
    console.log("Successully connected");
  };
});
