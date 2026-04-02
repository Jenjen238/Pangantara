### 11.1 Create Payment
**Definisi:** Membuat payment token Midtrans Snap untuk pesanan. Mengembalikan token dan redirect URL untuk menampilkan halaman pembayaran Midtrans.

**POST** `/payment/create`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `sppg`, `admin`

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_id | UUID | ✅ | ID pesanan |

**Contoh Request:**
```json
{
    "order_id": "uuid-pesanan"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Token payment berhasil dibuat |
| `400` | Validasi gagal atau pesanan tidak ditemukan |
| `401` | Token tidak valid |
| `403` | Role tidak diizinkan |
| `500` | Gagal membuat payment di Midtrans |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Payment created successfully",
    "data": {
        "token": "248eb2a6-622e-4828-97a4-23266c0d1899",
        "redirect_url": "https://app.sandbox.midtrans.com/snap/v4/redirection/248eb2a6-622e-4828-97a4-23266c0d1899"
    }
}