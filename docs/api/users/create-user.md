### 3.1 Create User
**Definisi:** Membuat user baru. Hanya dapat diakses oleh admin.

**POST** `/users`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ✅ | Nama lengkap |
| email | string | ✅ | Email valid dan unik |
| password | string | ✅ | Minimal 6 karakter |
| role | string | ✅ | `admin`, `supplier`, atau `sppg` |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | User berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |
