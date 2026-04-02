### 5.1 Create Supplier
**Definisi:** Membuat data supplier baru dengan status verifikasi `pending`.

**POST** `/suppliers`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user yang terkait |
| store_name | string | ✅ | Nama toko, maks 150 karakter |
| address | string | ❌ | Alamat toko |
| contact_number | string | ❌ | Nomor kontak, maks 20 karakter |
| category | string | ❌ | Kategori produk, maks 50 karakter |
| source_type | string | ❌ | Jenis sumber produk |
| business_desc | string | ❌ | Deskripsi bisnis |
| nib_document | string | ❌ | URL dokumen NIB |
| halal_document | string | ❌ | URL dokumen halal |
| other_document | string | ❌ | URL dokumen lainnya |
| admin_notes | string | ❌ | Catatan admin |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Supplier berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |