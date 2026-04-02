### 8.7 Delete Stock
**Definisi:** Menghapus data stok berdasarkan ID (soft delete).

**DELETE** `/stocks/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID stok |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Stok berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |
