import { searchInputHandler } from "./searchFilter.js";

const searchInput = document.querySelector("#search");

function eventHandlers() {
  searchInput.addEventListener("input", (e) => {
    searchInputHandler(e);
  });
}

export { eventHandlers };
