### 4.4 Get SPPG By User ID
**Definisi:** Mengambil data SPPG berdasarkan user ID.

**GET** `/sppg/user/:user_id`

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
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
