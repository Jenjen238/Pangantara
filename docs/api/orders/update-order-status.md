### 9.5 Update Order Status
**Definisi:** Memperbarui status pesanan. Hanya pesanan dengan status `pending` atau `cancelled` yang bisa dihapus.

**PUT** `/orders/:id/status`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID pesanan |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_status | string | ✅ | `pending`, `processing`, `shipped`, `completed`, `cancelled` |

**Contoh Request:**
```json
{
    "order_status": "shipped"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |