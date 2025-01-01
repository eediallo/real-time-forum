import { initiateChatHandler } from "./initiateChatHandler.js";
import { searchInputHandler } from "./seachHandler.js";

function eventHandlers() {
  searchInputHandler();
  initiateChatHandler();
}

export { eventHandlers };
