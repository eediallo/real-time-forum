async function fetchPrivateMessages() {
  try {
    const path = "/private_messages";
    const response = await fetch(path);
    if (!response.ok) {
      throw new Error("Failed to fetch private messages: ", response.status);
    }
    return await response.json();
  } catch (e) {
    console.log(e);
  }
}

export { fetchPrivateMessages };
