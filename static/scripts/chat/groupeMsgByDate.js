function groupMessagesByDate(messages) {
  const sortedMessages = messages.sort(
    (a, b) => new Date(a.createdAt) - new Date(b.createdAt)
  );

  return sortedMessages.reduce((acc, message) => {
    const date = new Date(message.createdAt).toLocaleDateString();
    if (!acc[date]) {
      acc[date] = [];
    }
    acc[date].push(message);
    return acc;
  }, {});
}

export { groupMessagesByDate };