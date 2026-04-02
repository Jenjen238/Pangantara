### 7.1 Create Product
**Definisi:** Membuat produk baru oleh supplier.

**POST** `/products`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| supplier_id | UUID | ✅ | ID supplier |
| product_name | string | ✅ | Nama produk, maks 150 karakter |
| category | string | ❌ | Kategori produk, maks 50 karakter |
| price | float64 | ✅ | Harga produk (harus lebih dari 0) |
| unit | string | ❌ | Satuan produk (kg, liter, dll) |
| image_url | string | ❌ | URL foto produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Produk berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |