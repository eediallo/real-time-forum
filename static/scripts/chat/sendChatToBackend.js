async function sendChatContent() {
  const chatInput = document.querySelector("#chatInput").value;
  const senderUsername = document.querySelector(".username").textContent;
  const receiverUsername = document.querySelector(".receiver-username").textContent;

  if (!chatInput) {
    alert("Message cannot be empty");
    return;
  }

  const payload = new URLSearchParams({
    chatInput: chatInput,
    senderUsername: senderUsername,
    receiverUsername: receiverUsername,
  });

  console.log("Payload:", payload.toString());

  try {
    const response = await fetch("http://localhost:8080/dashboard", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: payload,
    });

    if (response.ok) {
      const result = await response.text();
      console.log(result); // Message successfully sent.
    } else {
      const errorText = await response.text();
      console.error("Error:", response.statusText, errorText);
    }
  } catch (error) {
    console.error("Network error:", error);
  }
}

export { sendChatContent };