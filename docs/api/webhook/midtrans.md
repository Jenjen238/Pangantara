### 2.1 Midtrans Webhook
**Definisi:** Endpoint callback dari Midtrans untuk update status pembayaran secara otomatis.

**POST** `/webhook/midtrans`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| transaction_status | string | ✅ | Status dari Midtrans |
| order_id | string | ✅ | ID order |
| payment_type | string | ✅ | Metode pembayaran |
| fraud_status | string | ❌ | Status fraud |
| gross_amount | string | ✅ | Total pembayaran |
| signature_key | string | ✅ | Signature untuk verifikasi |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Notifikasi berhasil diproses |
| `400` | Signature tidak valid atau data tidak ditemukan 