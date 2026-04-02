### 8.5 Get Stock By Supplier ID
**Definisi:** Mengambil semua stok milik supplier tertentu.

**GET** `/stocks/supplier/:supplier_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| supplier_id | UUID (path) | ✅ | ID supplier |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
