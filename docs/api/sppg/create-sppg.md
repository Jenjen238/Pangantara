### 4.1 Create SPPG
**Definisi:** Membuat data SPPG baru.

**POST** `/sppg`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user yang terkait |
| name_sppg | string | ✅ | Nama SPPG, maks 150 karakter |
| location_url | string | ❌ | URL lokasi Google Maps |
| contact | string | ❌ | Nomor kontak, maks 20 karakter |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | SPPG berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |
