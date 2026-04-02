### 1.3 Refresh Token
**Definisi:** Memperbarui access token menggunakan refresh token yang masih valid.

**POST** `/auth/refresh`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| refresh_token | string | ✅ | Refresh token yang valid |

**Contoh Request:**
```json
{
    "refresh_token": "eyJhbGci..."
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Token berhasil diperbarui |
| `400` | Validasi gagal |
| `401` | Refresh token tidak valid atau expired |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Token berhasil diperbarui",
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci..."
}