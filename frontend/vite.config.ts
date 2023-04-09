import { defineConfig } from "vitest/config";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  test: {
    globals: true,
    environment: "jsdom",
    setupFiles: ["src/setupTest.js"],
    coverage: {
      provider: "c8",
      reporter: ["text", "html"],
    },
  },
});
