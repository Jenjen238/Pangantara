### 9.2 Get All Orders
**Definisi:** Mengambil semua pesanan dengan filter dan pagination.

**GET** `/orders`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter (Query):**
| Parameter | Type | Required | Default | Keterangan |
|-----------|------|----------|---------|------------|
| status | string | ❌ | - | Filter status: `pending`, `processing`, `shipped`, `completed`, `cancelled` |
| sppg_id | UUID | ❌ | - | Filter berdasarkan SPPG |
| start_date | string | ❌ | - | Filter tanggal mulai (YYYY-MM-DD) |
| end_date | string | ❌ | - | Filter tanggal akhir (YYYY-MM-DD) |
| page | int | ❌ | 1 | Halaman |
| limit | int | ❌ | 10 | Jumlah data per halaman |

**Contoh:**
```
GET /api/v1/orders?status=pending&page=1&limit=10
GET /api/v1/orders?start_date=2026-01-01&end_date=2026-03-31
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format parameter tidak valid |
| `401` | Token tidak valid |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Success",
    "data": [],
    "total": 50,
    "page": 1,
    "limit": 10
}