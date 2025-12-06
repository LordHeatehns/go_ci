import { defineConfig } from "@playwright/test";

export default defineConfig({
  testDir: "./tests", // โฟลเดอร์เก็บ test
  timeout: 10_000, // timeout ต่อ 1 test (ms)
  retries: 1, // retry เมื่อ fail (เหมาะสำหรับ CI)
  use: {
    baseURL: process.env.BASE_URL || "http://localhost:8855",
    extraHTTPHeaders: {
      "Content-Type": "application/json",
    },
  },
  reporter: [
    ["list"], // แสดงผลปกติใน CLI
    ["html", { outputFolder: "playwright-report" }], // generate report
  ],
});
