### 10.5 Update Payment Status
**Definisi:** Memperbarui status pembayaran. Jika `paid`, order status otomatis berubah ke `processing`. Jika `failed`, order status berubah ke `cancelled`. Email notifikasi dikirim ke SPPG.

**PUT** `/transactions/:id/status`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID transaksi |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| payment_status | string | ✅ | `unpaid`, `waiting_confirmation`, `paid`, `failed` |

**Contoh Request:**
```json
{
    "payment_status": "paid"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status pembayaran berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |