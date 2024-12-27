function validateForm(event) {
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    var errorMessage = '';

    if (!email) {
        errorMessage += 'Email is required.<br>';
    }
    if (!password) {
        errorMessage += 'Password is required.<br>';
    }

    if (errorMessage) {
        document.getElementById('error-message').innerHTML = errorMessage;
        event.preventDefault(); // Prevent form submission
    }
}