import { searchFilter } from "../searchFilter.js";

const searchInput = document.querySelector("#search");

function searchInputHandler() {
  searchInput.addEventListener("input", (e) => {
    searchFilter(e);
  });
}

export { searchInputHandler };
