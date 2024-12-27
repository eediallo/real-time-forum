function validateForm(event) {
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const nickName = document.getElementById("nickname").value;
    let errorMessage = "";

    if (!email && !nickName) {
        errorMessage += "Either email or nickname is required.<br>";
    }

    if (!password) {
        errorMessage += "Password is required.<br>";
    }

    if (errorMessage) {
        document.getElementById("error-message").innerHTML = errorMessage;
        event.preventDefault(); // Prevent form submission
    }
}
