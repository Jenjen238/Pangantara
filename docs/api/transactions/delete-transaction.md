### 10.6 Delete Transaction
**Definisi:** Menghapus transaksi berdasarkan ID (soft delete).

**DELETE** `/transactions/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID transaksi |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Transaksi berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |