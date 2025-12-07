//load/ — วัด performance ตาม usage จริง
import http from "k6/http";
import { sleep } from "k6";
import { loadTestOptions } from "../utils/options.js";

export const options = loadTestOptions();

export default function () {
  http.post("http://localhost:8855/api/v1/users/get/users/test");
  sleep(1);
}
