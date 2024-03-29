enum ColorThemes {
  Dark = "dark",
  Light = "light",
}

/**
 * checks if selected element or event target exists and is of type InputElement
 */
const isInputElement = (
  elem: HTMLElement | EventTarget | null
): elem is HTMLInputElement => !!elem && elem instanceof HTMLInputElement;

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
  isDark ? ColorThemes.Dark : ColorThemes.Light;

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
  e: Event,
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
      trackSystemColorModeChange(e, themeSwitcher, colorModeMediaQuery.matches);
    });

    const currentTheme = getStoredColorTheme();

    if (currentTheme === ColorThemes.Dark) {
      toggleThemeSwitcherState(themeSwitcher, true);
    }
    showIconLabel(document.querySelector(".toggle-switch__label"));
  }
};

export default () => {
  document.addEventListener("DOMContentLoaded", () => init());
};
