//stress/ — หาจุด “เริ่มพัง”
import http from "k6/http";
import { stressTestOptions } from "../utils/options.js";

export const options = stressTestOptions();

export default function () {
  http.post("http://localhost:8855/api/v1/users/get/users/test");
}
