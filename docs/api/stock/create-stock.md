### 8.1 Create Stock
**Definisi:** Membuat data stok untuk produk tertentu.

**POST** `/stocks`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| supplier_id | UUID | ✅ | ID supplier |
| product_id | UUID | ✅ | ID produk |
| stock_quantity | int | ✅ | Jumlah stok (minimal 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Stok berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |
