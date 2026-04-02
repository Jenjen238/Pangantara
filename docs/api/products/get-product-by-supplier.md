### 7.4 Get Product By Supplier
**Definisi:** Mengambil semua produk milik supplier tertentu, dengan filter kategori opsional.

**GET** `/products/supplier/:supplier_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| supplier_id | UUID (path) | ✅ | ID supplier |
| category | string (query) | ❌ | Filter kategori |

**Contoh:**
```
GET /api/v1/products/supplier/uuid?category=sayuran
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |