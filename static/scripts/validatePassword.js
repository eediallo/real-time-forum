function validatePassword(password) {
  if (password.length < 8) {
    throw new Error("Password must have at least 8 characters.");
  }

  if (!password.match(/[A-Z]/g)) {
    throw new Error("Password must have at least a capital letter.");
  }

  if (!password.match(/[0-9]/g)) {
    throw new Error("Password must have at least one digit.");
  }

  return password;
}

// const password = validatePassword("hellohelloA9");

console.log(password);

module.exports = validatePassword;
