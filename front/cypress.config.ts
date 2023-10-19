import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    baseUrl: "http://localhost:3002",
    env: {
      JWT_SECRET: "test",
    },
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
