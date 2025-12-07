import http from "k6/http";
import { spikeTestOptions } from "../utils/options.js";

export const options = spikeTestOptions();

export default function () {
  http.post("http://localhost:8855/api/v1/users/get/users/test");
}
