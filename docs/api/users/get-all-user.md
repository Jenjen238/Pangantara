
### 3.2 Get All Users
**Definisi:** Mengambil semua data user. Hanya dapat diakses oleh admin.

**GET** `/users`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
| `403` | Bukan admin |