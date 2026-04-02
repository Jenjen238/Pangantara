### 6.1 Save Draft
**Definisi:** Menyimpan progress pendaftaran supplier per step (create atau update otomatis).

**POST** `/supplier-draft/save`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user |
| store_name | string | ❌ | Nama toko |
| address | string | ❌ | Alamat |
| contact_number | string | ❌ | Nomor kontak |
| category | string | ❌ | Kategori produk |
| source_type | string | ❌ | Jenis sumber |
| business_desc | string | ❌ | Deskripsi bisnis |
| current_step | int | ❌ | Step form saat ini (1, 2, 3, dst) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Draft berhasil disimpan |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |