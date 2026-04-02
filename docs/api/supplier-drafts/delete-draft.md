### 6.5 Delete Draft
**Definisi:** Menghapus draft pendaftaran supplier berdasarkan user ID.

**DELETE** `/supplier-draft/:user_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Draft berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |