// Function to send chat content to the backend

async function sendChatContent() {
    const chatInput = document.querySelector("#chatInput").value;

    if (!chatInput) {
        alert("Message cannot be empty");
        return;
    }

    try {
        const response = await fetch("http://localhost:8080/dashboard", {
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: new URLSearchParams({ chatInput: chatInput }),
        });

        if (response.ok) {
            const result = await response.text();
            console.log(result); // Message successfully sent.
        } else {
            console.error("Error:", response.statusText);
        }
    } catch (error) {
        console.error("Network error:", error);
    }
}

export { sendChatContent };