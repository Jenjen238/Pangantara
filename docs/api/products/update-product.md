### 7.5 Update Product
**Definisi:** Memperbarui data produk berdasarkan ID.

**PUT** `/products/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| product_name | string | ❌ | Nama baru |
| category | string | ❌ | Kategori baru |
| price | float64 | ❌ | Harga baru (harus lebih dari 0) |
| unit | string | ❌ | Satuan baru |
| image_url | string | ❌ | URL foto baru |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Produk berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |