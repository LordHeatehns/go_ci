import { test, expect } from "@playwright/test";

test("POST /api/v1/users/get/users/test", async ({ request }) => {
  const res = await request.post("/api/v1/users/get/users/test", {
    data: {}, // ใส่ body ที่ระบบต้องการ ถ้าไม่มีปล่อยว่างได้
  });

  expect(res.status()).toBe(200);

  const body = await res.json();
  console.log(body);

  // ตัวอย่าง assertion
  expect(body).toBeDefined();
  // ถ้ารู้ structure ว่าต้อง return อะไร ใส่เพิ่มได้ เช่น:
  expect(body.data).toBeInstanceOf(Array);
});
