import Search from "./Search.svelte";
import initColorTheme from "./theme";

initColorTheme();

const target = document.getElementById("searchBlock");
const app = target ? new Search({ target }) : undefined;

export default app;
