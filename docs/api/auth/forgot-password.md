### 1.4 Forgot Password
**Definisi:** Mengirim email berisi link reset password ke email terdaftar.

**POST** `/auth/forgot-password`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| email | string | ✅ | Email terdaftar |

**Contoh Request:**
```json
{
    "email": "hanif@gmail.com"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Request berhasil (email dikirim jika terdaftar) |
| `400` | Validasi gagal |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "If your email is registered, a reset link will be sent",
    "data": null
}