### 6.2 Get Draft
**Definisi:** Mengambil data draft pendaftaran berdasarkan user ID.

**GET** `/supplier-draft/:user_id`

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
| `200` | Data draft berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Draft tidak ditemukan |
