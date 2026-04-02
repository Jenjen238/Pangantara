### 10.4 Get Transaction By Order ID
**Definisi:** Mengambil data transaksi berdasarkan order ID.

**GET** `/transactions/order/:order_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| order_id | UUID (path) | ✅ | ID pesanan |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Transaksi tidak ditemukan |