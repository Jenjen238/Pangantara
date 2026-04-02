### 7.2 Get All Products
**Definisi:** Mengambil semua produk dengan filter kategori opsional.

**GET** `/products`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter (Query):**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| category | string | ❌ | Filter berdasarkan kategori |

**Contoh:**
```
GET /api/v1/products?category=sayuran
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
