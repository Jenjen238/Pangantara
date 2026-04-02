### 1.5 Reset Password
**Definisi:** Mereset password menggunakan token dari email, berlaku 15 menit.

**POST** `/auth/reset-password`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| token | string | ✅ | Token dari link email |
| new_password | string | ✅ | Password baru minimal 6 karakter |
| confirm_password | string | ✅ | Harus sama dengan new_password |

**Contoh Request:**
```json
{
    "token": "abc123def456...",
    "new_password": "newpassword123",
    "confirm_password": "newpassword123"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Password berhasil direset |
| `400` | Token tidak valid, expired, atau password tidak cocok |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Password reset successfully",
    "data": null
}