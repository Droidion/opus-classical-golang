#!/usr/bin/env node

import esbuild from "esbuild";
import esbuildSvelte from "esbuild-svelte";
import sveltePreprocess from "svelte-preprocess";
import { compress } from "esbuild-plugin-compress";

esbuild
  .build({
    entryPoints: ["scripts/main.ts", "scripts/theme.ts"],
    mainFields: ["svelte", "browser", "module", "main"],
    bundle: true,
    outdir: "static/scripts",
    minify: true,
    write: false,
    plugins: [
      esbuildSvelte({
        preprocess: sveltePreprocess(),
      }),
      compress({
        outputDir: "",
      }),
    ],
    logLevel: "info",
  })
  .catch(() => process.exit(1));
