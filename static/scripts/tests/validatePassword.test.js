const validatePassword = require("../validatePassword.js");

describe("valide password", () => {
  test("Reject password of lenght less than 8", () => {
    const targetInput = "Password must have at least 8 characters.";

    expect(() => {
      validatePassword("hello");
    }).toThrow(targetInput);
  });

  test("Reject password which does not include at least a capital", () => {
    const targetInput = "Password must have at least a capital letter.";
    expect(() => {
      validatePassword("hellohello");
    }).toThrow(targetInput);
  });

  test("Reject password which does not include at least a digit", () => {
    const targetInput = "Password must have at least one digit.";

    expect(() => {
      validatePassword("hellAohello");
    }).toThrow(targetInput);
  });
});
