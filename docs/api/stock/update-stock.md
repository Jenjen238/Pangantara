### 8.6 Update Stock
**Definisi:** Memperbarui jumlah stok berdasarkan ID.

**PUT** `/stocks/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID stok |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| stock_quantity | int | ✅ | Jumlah stok baru (minimal 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Stok berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |a