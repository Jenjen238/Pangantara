### 9.6 Delete Order
**Definisi:** Menghapus pesanan berdasarkan ID. Hanya pesanan `pending` atau `cancelled` yang bisa dihapus.

**DELETE** `/orders/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID pesanan |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Pesanan berhasil dihapus |
| `400` | Pesanan tidak bisa dihapus karena statusnya |
| `401` | Token tidak valid |
| `404` | Pesanan tidak ditemukan |
| `500` | Server error |