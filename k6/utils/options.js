//utils/ — ไฟล์ shared เช่น auth / headers / options
export function smokeTestOptions() {
  return {
    vus: 1,
    duration: "10s",
  };
}

export function loadTestOptions() {
  return {
    vus: 20,
    duration: "30s",
  };
}

export function stressTestOptions() {
  return {
    stages: [
      { duration: "30s", target: 20 },
      { duration: "30s", target: 50 },
      { duration: "30s", target: 0 },
    ],
  };
}

export function soakTestOptions() {
  return {
    vus: 10,
    duration: "10m",
  };
}

export function spikeTestOptions() {
  return {
    stages: [
      { duration: "10s", target: 2 },
      { duration: "2s", target: 200 },
      { duration: "30s", target: 5 },
    ],
  };
}
