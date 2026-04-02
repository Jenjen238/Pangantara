### 1.1 Register
**Definisi:** Mendaftarkan user baru ke platform Pangantara.

**POST** `/auth/register`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ✅ | Nama lengkap user |
| email | string | ✅ | Email valid |
| password | string | ✅ | Minimal 6 karakter |
| role | string | ✅ | `admin`, `supplier`, atau `sppg` |

**Contoh Request:**
```json
{
    "name": "Hanif Afghani",
    "email": "hanif@gmail.com",
    "password": "123456",
    "role": "admin"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Register berhasil |
| `400` | Validasi gagal atau email sudah digunakan |

**Contoh Response (201):**
```json
{
    "success": true,
    "message": "Created successfully",
    "data": {
        "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6",
        "name": "Hanif Afghani",
        "email": "hanif@gmail.com",
        "role": "admin",
        "created_at": "2026-03-25T00:00:00Z",
        "updated_at": "2026-03-25T00:00:00Z"
    }
}