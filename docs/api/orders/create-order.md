### 9.1 Create Order
**Definisi:** Membuat pesanan baru oleh SPPG dengan multiple item. Total harga dihitung otomatis.

**POST** `/orders`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| sppg_id | UUID | ✅ | ID SPPG |
| notes | string | ❌ | Catatan pesanan |
| items | array | ✅ | Minimal 1 item |
| items[].product_id | UUID | ✅ | ID produk |
| items[].quantity | int | ✅ | Jumlah (harus lebih dari 0) |

**Contoh Request:**
```json
{
    "sppg_id": "uuid-sppg",
    "notes": "Tolong kirim pagi hari",
    "items": [
        {
            "product_id": "uuid-produk-1",
            "quantity": 5
        },
        {
            "product_id": "uuid-produk-2",
            "quantity": 10
        }
    ]
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Pesanan berhasil dibuat |
| `400` | Validasi gagal atau produk tidak ditemukan |
| `401` | Token tidak valid |
| `500` | Server error |