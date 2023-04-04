module.exports = {
  "env": {
    "es6": true,
    "browser": true
  },
  "plugins": ["prettier", "@typescript-eslint"],
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "prettier",
    "plugin:svelte/recommended"
  ],
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "parser": "@typescript-eslint/parser",
    "sourceType": "module",
    "ecmaVersion": 2019,
    "extraFileExtensions": [".svelte"],
    "project": "./tsconfig.json"
  },
  "rules": {
    "prettier/prettier": ["error"]
  },
  "root": true,
  "overrides": [
    {
      "files": ["*.svelte"],
      "parser": "svelte-eslint-parser",
      "parserOptions": {
        "parser": "@typescript-eslint/parser",
      },
    },
  ],
}