import Search from "./Search.svelte";
import initColorTheme from "./theme";

initColorTheme();

const target = document.getElementById("searchBlock");

if (target) {
  new Search({ target });
}
