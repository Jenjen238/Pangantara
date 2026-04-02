### 5.6 Verify Supplier
**Definisi:** Approve atau reject supplier oleh admin. Akan mengirim email notifikasi ke supplier.

**PATCH** `/suppliers/:id/verify`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| status | string | ✅ | `approved` atau `rejected` |
| admin_notes | string | ❌ | Alasan approve/reject |

**Contoh Request:**
```json
{
    "status": "approved",
    "admin_notes": "Dokumen lengkap dan valid"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status verifikasi berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |
