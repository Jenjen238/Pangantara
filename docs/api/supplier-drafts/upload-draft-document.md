### 6.4 Upload Draft Document
**Definisi:** Upload dokumen (NIB, Halal, Other) untuk draft pendaftaran supplier.

**PATCH** `/supplier-draft/:user_id/document`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | multipart/form-data | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Form Data:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| document_type | string | ✅ | `nib`, `halal`, atau `other` |
| file | file | ✅ | File PDF, JPG, atau PNG, maks 5MB |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Dokumen berhasil diupload |
| `400` | Tipe file tidak valid atau ukuran melebihi batas |
| `401` | Token tidak valid |
| `500` | Server error |