### 7.6 Delete Product
**Definisi:** Menghapus produk berdasarkan ID (soft delete).

**DELETE** `/products/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Produk berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |
