### 9.4 Get Order By SPPG ID
**Definisi:** Mengambil semua pesanan milik SPPG tertentu.

**GET** `/orders/sppg/:sppg_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| sppg_id | UUID (path) | ✅ | ID SPPG |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |