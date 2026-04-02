### 6.3 Submit Draft
**Definisi:** Submit pendaftaran supplier dari draft menjadi supplier dengan status `pending`.

**POST** `/supplier-draft/submit`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user |

**Contoh Request:**
```json
{
    "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Pendaftaran berhasil disubmit |
| `400` | Draft tidak ditemukan atau field wajib belum diisi |
| `401` | Token tidak valid |
| `500` | Server error |