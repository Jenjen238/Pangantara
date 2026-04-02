### 4.5 Update SPPG
**Definisi:** Memperbarui data SPPG berdasarkan ID.

**PUT** `/sppg/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID SPPG |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name_sppg | string | ❌ | Nama baru SPPG |
| location_url | string | ❌ | URL lokasi baru |
| contact | string | ❌ | Nomor kontak baru |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | SPPG berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |
