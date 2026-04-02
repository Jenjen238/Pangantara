### 5.5 Update Supplier
**Definisi:** Memperbarui data supplier berdasarkan ID.

**PUT** `/suppliers/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| store_name | string | ❌ | Nama toko baru |
| address | string | ❌ | Alamat baru |
| contact_number | string | ❌ | Nomor kontak baru |
| category | string | ❌ | Kategori baru |
| source_type | string | ❌ | Jenis sumber baru |
| business_desc | string | ❌ | Deskripsi bisnis baru |
| nib_document | string | ❌ | URL dokumen NIB baru |
| halal_document | string | ❌ | URL dokumen halal baru |
| other_document | string | ❌ | URL dokumen lain baru |
| admin_notes | string | ❌ | Catatan admin baru |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Supplier berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |