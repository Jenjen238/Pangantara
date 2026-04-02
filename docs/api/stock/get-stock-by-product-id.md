### 8.4 Get Stock By Product ID
**Definisi:** Mengambil data stok berdasarkan product ID.

**GET** `/stocks/product/:product_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| product_id | UUID (path) | ✅ | ID produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Stok tidak ditemukan |
