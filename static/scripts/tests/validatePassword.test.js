const validatePassword = require("../validatePassword.js");

describe("valide password", () => {
  test("Reject password of lenght less than 8", () => {
    const targetInput = "Password does not meet mininum requirement.";

    expect(() => {
      validatePassword("hello");
    }).toThrow(targetInput);
  });
});
