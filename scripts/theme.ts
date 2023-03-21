/**
 * sets data-attribute 'theme' of the root html element
 * @param colorTheme theme name
 */
const setColorTheme = (colorTheme: string): void => {
  document.documentElement.dataset.theme = colorTheme;
};

/**
 * stores current color theme in the localStorage
 * @param colorTheme theme name
 */
const storeColorTheme = (colorTheme: string): void => {
  localStorage.setItem("theme", colorTheme);
};

/**
 * gets previously stored color theme from the localStorage
 * @returns color theme name if any
 */
const getStoredColorTheme = (): string | null => {
  const storedColorTheme = localStorage.getItem("theme");
  return storedColorTheme;
};

/**
 * toggles dark/light color mode
 * @param isDark if checkbox checked
 */
const toggleColorTheme = (isDark: boolean): string =>
  isDark ? "dark" : "light";

/**
 * defines init color theme in the following priority order: 1. prev stored color theme, 2. system color scheme, 3. theme light;
 * adds event listener for tracking theme switch and system color scheme change
 */
const defineColorThemeOnLoad = () => {
  const themeSwitcher = document.getElementById(
    "switcher"
  )! as HTMLInputElement;

  const colorModeMediaQuery = window.matchMedia("(prefers-color-scheme: dark)");

  themeSwitcher.addEventListener("change", (e) => {
    const colorTheme = toggleColorTheme(
      (e!.target as HTMLInputElement).checked
    );
    storeColorTheme(colorTheme);
    setColorTheme(colorTheme);
    console.log("toggle currentTheme", colorTheme);
  });

  colorModeMediaQuery.addEventListener("change", () => {
    const colorTheme = toggleColorTheme(colorModeMediaQuery.matches);
    storeColorTheme(colorTheme);
    setColorTheme(colorTheme);
    themeSwitcher.checked = colorTheme === "dark";
    console.log("system change currentTheme", colorTheme);
  });

  let currentTheme = "light";

  const isSystemColorThemeDark = colorModeMediaQuery.matches;
  console.log("isSystemColorThemeDark", isSystemColorThemeDark);
  const prevStoredTheme = getStoredColorTheme();
  const prevStoredDark = prevStoredTheme && prevStoredTheme === "dark";

  if (prevStoredDark || (!prevStoredTheme && isSystemColorThemeDark)) {
    currentTheme = "dark";
    themeSwitcher.checked = true;
  }

  setColorTheme(currentTheme);

  console.log("currentTheme on load", currentTheme);
};

export default (() => {
  document.addEventListener("DOMContentLoaded", () => defineColorThemeOnLoad());
})();
