### 4.2 Get All SPPG
**Definisi:** Mengambil semua data SPPG.

**GET** `/sppg`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
