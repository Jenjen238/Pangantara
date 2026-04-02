### 9.3 Get Order By ID
**Definisi:** Mengambil detail pesanan berdasarkan ID, termasuk detail item dan transaksi.

**GET** `/orders/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID pesanan |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Pesanan tidak ditemukan |
