### 10.2 Get All Transactions
**Definisi:** Mengambil semua data transaksi.

**GET** `/transactions`

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
