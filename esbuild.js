#!/usr/bin/env node

import esbuild from "esbuild";
import esbuildSvelte from "esbuild-svelte";
import sveltePreprocess from "svelte-preprocess";
import { compress } from 'esbuild-plugin-compress';
import {sassPlugin} from 'esbuild-sass-plugin';

esbuild
    .build({
        entryPoints: ["frontend/scripts/main.ts", "frontend/styles/site.sass"],
        mainFields: ["svelte", "browser", "module", "main"],
        bundle: true,
        outdir: "static",
        minify: true,
        write: false,
        plugins: [
            esbuildSvelte({
                preprocess: sveltePreprocess(),
            }),
            compress({
                outputDir: '',
            }),
            sassPlugin()
        ],
        logLevel: "info",
    })
    .catch(() => process.exit(1))