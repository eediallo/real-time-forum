import { validPassword } from "./validePassword.js";
const errorMsgEl = document.querySelector("#error-message");
const submitBtn = document.querySelector("#signUpBtn");

function validatePassword(event) {
  const password = document.querySelector("#password").value;
  let errorMessage =
    "The password you entered does not meet our security requirements. Please ensure it includes at least 8 characters, an uppercase letter, a number, and a special character.";

  try {
    if (!validPassword(password)) {
      event.preventDefault();
      errorMsgEl.innerHTML = errorMessage;
      errorMsgEl.style.color = "red";
    } else {
      errorMsgEl.innerHTML = ""; // Clear the error message if the password is valid
      console.log("Password Accepted.");
    }
  } catch (error) {
    // Use the thrown error message for UI feedback
    console.error("invalid password", error.message);
    event.preventDefault();
    errorMsgEl.innerHTML = error.message;
    errorMsgEl.style.color = "red";
  }
}

function validePasswordHandler() {
  submitBtn.addEventListener("click", (e) => {
    validatePassword(e);
  });
}

validePasswordHandler();
