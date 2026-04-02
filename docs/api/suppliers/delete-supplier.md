### 5.7 Delete Supplier
**Definisi:** Menghapus data supplier berdasarkan ID (soft delete).

**DELETE** `/suppliers/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Supplier berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |