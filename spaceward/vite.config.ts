import path from "path";
import { defineConfig } from "vitest/config";
import { nodePolyfills } from "vite-plugin-node-polyfills";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
    server: {
        proxy: {
            "/api/faucet": {
                target: "http://localhost:8000/",
                rewrite: (path) => path.replace(/^\/api\/faucet/, "/"),
            },
        },
    },
    plugins: [
        react(),
        nodePolyfills({
            // To exclude specific polyfills, add them to this list.
            exclude: [
                "fs", // Excludes the polyfill for `fs` and `node:fs`.
            ],
            // Whether to polyfill specific globals.
            globals: {
                Buffer: true, // can also be 'build', 'dev', or false
                global: true,
                process: false,
            },
            // Whether to polyfill `node:` protocol imports.
            protocolImports: false,
        }),
    ],
    test: {
        environment: "jsdom",
        globals: true,
        setupFiles: "./tests/setup.ts", // assuming the test folder is in the root of our project
    },
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
        },
    },
});
