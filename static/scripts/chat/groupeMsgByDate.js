function groupMessagesByDate(messages) {
  return messages.reduce((acc, message) => {
    const date = new Date(message.createdAt).toLocaleDateString();
    if (!acc[date]) {
      acc[date] = [];
    }
    acc[date].push(message);
    return acc;
  }, {});
}

export { groupMessagesByDate };
