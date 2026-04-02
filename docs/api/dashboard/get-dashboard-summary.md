### 13.1 Get Dashboard Summary
**Definisi:** Mengambil ringkasan data untuk dashboard admin.

**GET** `/dashboard/summary`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Success",
    "data": {
        "total_supplier": 50,
        "supplier_pending": 10,
        "supplier_approved": 35,
        "supplier_rejected": 5,
        "total_sppg": 120,
        "total_order": 500,
        "order_pending": 20,
        "order_processing": 15,
        "order_shipped": 10,
        "order_completed": 450,
        "order_cancelled": 5
    }
}