### 3.4 Update User
**Definisi:** Memperbarui data user berdasarkan ID.

**PUT** `/users/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`, `supplier`, `sppg`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID user |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ❌ | Nama baru |
| email | string | ❌ | Email baru (harus valid) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | User berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Role tidak diizinkan |
| `500` | Server error |
