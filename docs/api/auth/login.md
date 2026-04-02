### 1.2 Login
**Definisi:** Autentikasi user dan mendapatkan access token & refresh token.

**POST** `/auth/login`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| email | string | ✅ | Email terdaftar |
| password | string | ✅ | Password user |

**Contoh Request:**
```json
{
    "email": "hanif@gmail.com",
    "password": "123456"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Login berhasil |
| `400` | Validasi gagal |
| `401` | Email atau password salah |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Login berhasil",
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci...",
    "data": {
        "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6",
        "name": "Hanif Afghani",
        "email": "hanif@gmail.com",
        "role": "admin"
    }
}