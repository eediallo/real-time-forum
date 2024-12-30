const addChars = /['!""#$%&"()*,-./]/g;

function validPassword(password) {
  if (typeof password !== "string") {
    throw new Error("Password must be a string.");
  }

  if (password.length < 8) {
    throw new Error("Password must have at least 8 characters.");
  }

  if (!password.match(/[A-Z]/g)) {
    throw new Error("Password must have at least a capital letter.");
  }

  if (!password.match(/[0-9]/g)) {
    throw new Error("Password must have at least one digit.");
  }

  if (!password.match(addChars)) {
    throw new Error(
      "Password must have at least least one additional character."
    );
  }

  return true;
}

export { validPassword };
//module.exports = validPassword;
