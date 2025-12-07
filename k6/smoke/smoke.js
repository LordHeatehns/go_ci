import http from "k6/http";
import { sleep } from "k6";
import { smokeTestOptions } from "../utils/options.js";

export const options = smokeTestOptions();

export default function () {
  http.get("http://localhost:8855/health");
  sleep(1);
}
