### 4.6 Delete SPPG
**Definisi:** Menghapus data SPPG berdasarkan ID (soft delete).

**DELETE** `/sppg/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID SPPG |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | SPPG berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |
