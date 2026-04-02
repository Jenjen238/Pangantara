### 8.2 Get All Stocks
**Definisi:** Mengambil semua data stok.

**GET** `/stocks`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
