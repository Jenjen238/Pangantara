### 5.2 Get All Suppliers
**Definisi:** Mengambil semua data supplier dengan filter opsional.

**GET** `/suppliers`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter (Query):**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| keyword | string | ❌ | Cari berdasarkan nama toko |
| category | string | ❌ | Filter berdasarkan kategori |
| status | string | ❌ | Filter: `pending`, `approved`, `rejected` |

**Contoh:**
```
GET /api/v1/suppliers?keyword=toko&category=sayuran&status=approved
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
