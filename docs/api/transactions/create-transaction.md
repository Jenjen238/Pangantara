### 10.1 Create Transaction
**Definisi:** Membuat transaksi pembayaran untuk pesanan tertentu.

**POST** `/transactions`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_id | UUID | ✅ | ID pesanan |
| payment_method | string | ❌ | Metode pembayaran (transfer, dll) |
| payment_proof | string | ❌ | URL bukti transfer |
| amount_paid | float64 | ✅ | Jumlah yang dibayar (harus lebih dari 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Transaksi berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |