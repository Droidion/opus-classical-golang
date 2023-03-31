/**
 * checks if selected element or event target exists and is of type InputElement
 */
const isInputElement = (
  elem: HTMLElement | EventTarget | null
): elem is HTMLInputElement =>
  Boolean(elem) && elem instanceof HTMLInputElement;

/**
 * sets data-attribute 'theme' of the root html element
 */
const setColorTheme = (themeName: string): void => {
  document.documentElement.dataset.theme = themeName;
};

/**
 * stores current color theme in the localStorage
 */
const storeColorTheme = (themeName: string): void => {
  localStorage.setItem("theme", themeName);
};

/**
 * gets previously stored color theme from the localStorage
 */
const getStoredColorTheme = (): string | null => localStorage.getItem("theme");

/**
 * toggles dark/light color theme
 */
const toggleColorTheme = (isDark: boolean): string =>
  isDark ? "dark" : "light";

/**
 *  toggles checkbox checked state
 */
const toggleThemeSwitcherState = (
  switcher: HTMLElement,
  isDark: boolean
): void => {
  if (isInputElement(switcher)) {
    switcher.checked = isDark;
  }
};

/**
 * cb triggered when checkbox is checked
 */
const trackColorThemeChange = ({ target }: Event): void => {
  if (isInputElement(target)) {
    const colorTheme = toggleColorTheme(target.checked);
    storeColorTheme(colorTheme);
    setColorTheme(colorTheme);
  }
};

/**
 * cb triggered when user changes system color scheme
 */
const trackSystemColorModeChange = (
  switcher: HTMLElement,
  isSystemDarkModeOn: boolean
): void => {
  const colorTheme = toggleColorTheme(isSystemDarkModeOn);

  storeColorTheme(colorTheme);
  setColorTheme(colorTheme);
  toggleThemeSwitcherState(switcher, isSystemDarkModeOn);
};

/**
 * shows checkbox label (sun/moon icon)
 */
const showIconLabel = (label: HTMLLabelElement | null): void => {
  if (label) {
    label.classList.remove("d-none");
  }
};

/**
 * defines current color theme in the following priority order:
 * 1. prev stored color theme, 2. system color scheme, 3. theme light;
 */
const defineColorTheme = (isSystemColorThemeDark: boolean): string => {
  let currentTheme = "light";

  const prevStoredTheme = getStoredColorTheme();
  const isPrevStoredDark = Boolean(
    prevStoredTheme && prevStoredTheme === "dark"
  );

  if (isPrevStoredDark || (!prevStoredTheme && isSystemColorThemeDark)) {
    currentTheme = "dark";
  }

  return currentTheme;
};

/**
 * defines color theme and checkbox state;
 * adds event listener for tracking theme switch and system color scheme change;
 * makes theme switcher label visible
 */
const init = () => {
  const themeSwitcher = document.getElementById("switcher");

  const colorModeMediaQuery = window.matchMedia("(prefers-color-scheme: dark)");

  if (themeSwitcher) {
    themeSwitcher.addEventListener("change", trackColorThemeChange);

    colorModeMediaQuery.addEventListener("change", (e) => {
      trackSystemColorModeChange(themeSwitcher, colorModeMediaQuery.matches);
    });

    const currentTheme = defineColorTheme(colorModeMediaQuery.matches);

    toggleThemeSwitcherState(themeSwitcher, currentTheme === "dark");
    setColorTheme(currentTheme);
    showIconLabel(document.querySelector(".toggle-switch__label"));
  }
};

export default () => {
  document.addEventListener("DOMContentLoaded", () => init());
};
