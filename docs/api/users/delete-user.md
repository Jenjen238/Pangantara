### 3.5 Delete User
**Definisi:** Menghapus user berdasarkan ID (soft delete).

**DELETE** `/users/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | User berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |