# API Documentation - Pangantara
**Version:** 1.0.0
**Base URL:** `http://localhost:8080/api/v1`
**Content-Type:** `application/json`

---

## 📌 Definisi

| Term | Definisi |
|------|----------|
| `access_token` | JWT token yang digunakan untuk autentikasi, berlaku **15 menit** |
| `refresh_token` | JWT token untuk memperbarui access token, berlaku **7 hari** |
| `UUID` | Unique identifier berbentuk string, contoh: `f00a42ce-60e7-4170-8eaa-13c4776603b6` |
| `Bearer Token` | Format token di header: `Authorization: Bearer <access_token>` |
| `Role` | Hak akses pengguna: `admin`, `supplier`, `sppg` |
| `Verification Status` | Status verifikasi supplier: `pending`, `approved`, `rejected` |
| `Order Status` | Status pesanan: `pending`, `processing`, `shipped`, `completed`, `cancelled` |
| `Payment Status` | Status pembayaran: `unpaid`, `waiting_confirmation`, `paid`, `failed` |
| `Draft Status` | Status draft pendaftaran supplier: `draft`, `submitted` |

---

## 📌 Format Response Standar

### Success Response
```json
{
    "success": true,
    "message": "Success",
    "data": {}
}
```

### Paginated Response
```json
{
    "success": true,
    "message": "Success",
    "data": [],
    "total": 100,
    "page": 1,
    "limit": 10
}
```

### Error Response
```json
{
    "success": false,
    "message": "Error message"
}
```

---

## 📌 HTTP Status Code

| Status Code | Keterangan |
|-------------|------------|
| `200` | OK - Request berhasil |
| `201` | Created - Data berhasil dibuat |
| `400` | Bad Request - Request tidak valid |
| `401` | Unauthorized - Token tidak ada atau tidak valid |
| `403` | Forbidden - Tidak punya izin akses |
| `404` | Not Found - Data tidak ditemukan |
| `429` | Too Many Requests - Rate limit tercapai |
| `500` | Internal Server Error - Kesalahan server |

---

## 📌 Rate Limiting

| Endpoint | Limit |
|----------|-------|
| Global (semua endpoint) | 100 request/menit |
| Auth (login, register, forgot password) | 10 request/menit |
| Upload (dokumen & foto) | 20 request/menit |

---

# 🔓 PUBLIC ENDPOINTS

---

## 1. Auth

### 1.1 Register
**Definisi:** Mendaftarkan user baru ke platform Pangantara.

**POST** `/auth/register`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ✅ | Nama lengkap user |
| email | string | ✅ | Email valid |
| password | string | ✅ | Minimal 6 karakter |
| role | string | ✅ | `admin`, `supplier`, atau `sppg` |

**Contoh Request:**
```json
{
    "name": "Hanif Afghani",
    "email": "hanif@gmail.com",
    "password": "123456",
    "role": "admin"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Register berhasil |
| `400` | Validasi gagal atau email sudah digunakan |

**Contoh Response (201):**
```json
{
    "success": true,
    "message": "Created successfully",
    "data": {
        "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6",
        "name": "Hanif Afghani",
        "email": "hanif@gmail.com",
        "role": "admin",
        "created_at": "2026-03-25T00:00:00Z",
        "updated_at": "2026-03-25T00:00:00Z"
    }
}
```

---

### 1.2 Login
**Definisi:** Autentikasi user dan mendapatkan access token & refresh token.

**POST** `/auth/login`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| email | string | ✅ | Email terdaftar |
| password | string | ✅ | Password user |

**Contoh Request:**
```json
{
    "email": "hanif@gmail.com",
    "password": "123456"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Login berhasil |
| `400` | Validasi gagal |
| `401` | Email atau password salah |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Login berhasil",
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci...",
    "data": {
        "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6",
        "name": "Hanif Afghani",
        "email": "hanif@gmail.com",
        "role": "admin"
    }
}
```

---

### 1.3 Refresh Token
**Definisi:** Memperbarui access token menggunakan refresh token yang masih valid.

**POST** `/auth/refresh`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| refresh_token | string | ✅ | Refresh token yang valid |

**Contoh Request:**
```json
{
    "refresh_token": "eyJhbGci..."
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Token berhasil diperbarui |
| `400` | Validasi gagal |
| `401` | Refresh token tidak valid atau expired |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Token berhasil diperbarui",
    "access_token": "eyJhbGci...",
    "refresh_token": "eyJhbGci..."
}
```

---

### 1.4 Forgot Password
**Definisi:** Mengirim email berisi link reset password ke email terdaftar.

**POST** `/auth/forgot-password`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| email | string | ✅ | Email terdaftar |

**Contoh Request:**
```json
{
    "email": "hanif@gmail.com"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Request berhasil (email dikirim jika terdaftar) |
| `400` | Validasi gagal |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "If your email is registered, a reset link will be sent",
    "data": null
}
```

---

### 1.5 Reset Password
**Definisi:** Mereset password menggunakan token dari email, berlaku 15 menit.

**POST** `/auth/reset-password`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| token | string | ✅ | Token dari link email |
| new_password | string | ✅ | Password baru minimal 6 karakter |
| confirm_password | string | ✅ | Harus sama dengan new_password |

**Contoh Request:**
```json
{
    "token": "abc123def456...",
    "new_password": "newpassword123",
    "confirm_password": "newpassword123"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Password berhasil direset |
| `400` | Token tidak valid, expired, atau password tidak cocok |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Password reset successfully",
    "data": null
}
```

---

## 2. Webhook

### 2.1 Midtrans Webhook
**Definisi:** Endpoint callback dari Midtrans untuk update status pembayaran secara otomatis.

**POST** `/webhook/midtrans`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| transaction_status | string | ✅ | Status dari Midtrans |
| order_id | string | ✅ | ID order |
| payment_type | string | ✅ | Metode pembayaran |
| fraud_status | string | ❌ | Status fraud |
| gross_amount | string | ✅ | Total pembayaran |
| signature_key | string | ✅ | Signature untuk verifikasi |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Notifikasi berhasil diproses |
| `400` | Signature tidak valid atau data tidak ditemukan |

---

# 🔐 PROTECTED ENDPOINTS

> Semua endpoint di bawah memerlukan header:
> ```
> Authorization: Bearer <access_token>
> ```

---

## 3. Users

### 3.1 Create User
**Definisi:** Membuat user baru. Hanya dapat diakses oleh admin.

**POST** `/users`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ✅ | Nama lengkap |
| email | string | ✅ | Email valid dan unik |
| password | string | ✅ | Minimal 6 karakter |
| role | string | ✅ | `admin`, `supplier`, atau `sppg` |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | User berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

---

### 3.2 Get All Users
**Definisi:** Mengambil semua data user. Hanya dapat diakses oleh admin.

**GET** `/users`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
| `403` | Bukan admin |

---

### 3.3 Get User By ID
**Definisi:** Mengambil data user berdasarkan ID.

**GET** `/users/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`, `supplier`, `sppg`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `403` | Role tidak diizinkan |
| `404` | User tidak ditemukan |

---

### 3.4 Update User
**Definisi:** Memperbarui data user berdasarkan ID.

**PUT** `/users/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`, `supplier`, `sppg`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID user |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name | string | ❌ | Nama baru |
| email | string | ❌ | Email baru (harus valid) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | User berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Role tidak diizinkan |
| `500` | Server error |

---

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

---

## 4. SPPG

### 4.1 Create SPPG
**Definisi:** Membuat data SPPG baru.

**POST** `/sppg`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user yang terkait |
| name_sppg | string | ✅ | Nama SPPG, maks 150 karakter |
| location_url | string | ❌ | URL lokasi Google Maps |
| contact | string | ❌ | Nomor kontak, maks 20 karakter |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | SPPG berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 4.2 Get All SPPG
**Definisi:** Mengambil semua data SPPG.

**GET** `/sppg`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |

---

### 4.3 Get SPPG By ID
**Definisi:** Mengambil data SPPG berdasarkan ID.

**GET** `/sppg/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID SPPG |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | SPPG tidak ditemukan |

---

### 4.4 Get SPPG By User ID
**Definisi:** Mengambil data SPPG berdasarkan user ID.

**GET** `/sppg/user/:user_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |

---

### 4.5 Update SPPG
**Definisi:** Memperbarui data SPPG berdasarkan ID.

**PUT** `/sppg/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID SPPG |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| name_sppg | string | ❌ | Nama baru SPPG |
| location_url | string | ❌ | URL lokasi baru |
| contact | string | ❌ | Nomor kontak baru |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | SPPG berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 4.6 Delete SPPG
**Definisi:** Menghapus data SPPG berdasarkan ID (soft delete).

**DELETE** `/sppg/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID SPPG |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | SPPG berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |

---

## 5. Suppliers

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

---

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

---

### 5.3 Get Supplier By ID
**Definisi:** Mengambil data supplier berdasarkan ID.

**GET** `/suppliers/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`, `supplier`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Supplier tidak ditemukan |

---

### 5.4 Get Supplier By User ID
**Definisi:** Mengambil data supplier berdasarkan user ID.

**GET** `/suppliers/user/:user_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Supplier tidak ditemukan |

---

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

---

### 5.6 Verify Supplier
**Definisi:** Approve atau reject supplier oleh admin. Akan mengirim email notifikasi ke supplier.

**PATCH** `/suppliers/:id/verify`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| status | string | ✅ | `approved` atau `rejected` |
| admin_notes | string | ❌ | Alasan approve/reject |

**Contoh Request:**
```json
{
    "status": "approved",
    "admin_notes": "Dokumen lengkap dan valid"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status verifikasi berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

---

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

---

## 6. Supplier Draft

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

---

### 6.2 Get Draft
**Definisi:** Mengambil data draft pendaftaran berdasarkan user ID.

**GET** `/supplier-draft/:user_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data draft berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Draft tidak ditemukan |

---

### 6.3 Submit Draft
**Definisi:** Submit pendaftaran supplier dari draft menjadi supplier dengan status `pending`.

**POST** `/supplier-draft/submit`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| user_id | UUID | ✅ | ID user |

**Contoh Request:**
```json
{
    "user_id": "f00a42ce-60e7-4170-8eaa-13c4776603b6"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Pendaftaran berhasil disubmit |
| `400` | Draft tidak ditemukan atau field wajib belum diisi |
| `401` | Token tidak valid |
| `500` | Server error |

---

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

---

### 6.5 Delete Draft
**Definisi:** Menghapus draft pendaftaran supplier berdasarkan user ID.

**DELETE** `/supplier-draft/:user_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| user_id | UUID (path) | ✅ | ID user |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Draft berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |

---

## 7. Products

### 7.1 Create Product
**Definisi:** Membuat produk baru oleh supplier.

**POST** `/products`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| supplier_id | UUID | ✅ | ID supplier |
| product_name | string | ✅ | Nama produk, maks 150 karakter |
| category | string | ❌ | Kategori produk, maks 50 karakter |
| price | float64 | ✅ | Harga produk (harus lebih dari 0) |
| unit | string | ❌ | Satuan produk (kg, liter, dll) |
| image_url | string | ❌ | URL foto produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Produk berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 7.2 Get All Products
**Definisi:** Mengambil semua produk dengan filter kategori opsional.

**GET** `/products`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter (Query):**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| category | string | ❌ | Filter berdasarkan kategori |

**Contoh:**
```
GET /api/v1/products?category=sayuran
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |

---

### 7.3 Get Product By ID
**Definisi:** Mengambil data produk berdasarkan ID.

**GET** `/products/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Produk tidak ditemukan |

---

### 7.4 Get Product By Supplier
**Definisi:** Mengambil semua produk milik supplier tertentu, dengan filter kategori opsional.

**GET** `/products/supplier/:supplier_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| supplier_id | UUID (path) | ✅ | ID supplier |
| category | string (query) | ❌ | Filter kategori |

**Contoh:**
```
GET /api/v1/products/supplier/uuid?category=sayuran
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |

---

### 7.5 Update Product
**Definisi:** Memperbarui data produk berdasarkan ID.

**PUT** `/products/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| product_name | string | ❌ | Nama baru |
| category | string | ❌ | Kategori baru |
| price | float64 | ❌ | Harga baru (harus lebih dari 0) |
| unit | string | ❌ | Satuan baru |
| image_url | string | ❌ | URL foto baru |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Produk berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 7.6 Delete Product
**Definisi:** Menghapus produk berdasarkan ID (soft delete).

**DELETE** `/products/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Produk berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |

---

## 8. Stocks

### 8.1 Create Stock
**Definisi:** Membuat data stok untuk produk tertentu.

**POST** `/stocks`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| supplier_id | UUID | ✅ | ID supplier |
| product_id | UUID | ✅ | ID produk |
| stock_quantity | int | ✅ | Jumlah stok (minimal 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Stok berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 8.2 Get All Stocks
**Definisi:** Mengambil semua data stok.

**GET** `/stocks`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |

---

### 8.3 Get Stock By ID
**Definisi:** Mengambil data stok berdasarkan ID.

**GET** `/stocks/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID stok |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Stok tidak ditemukan |

---

### 8.4 Get Stock By Product ID
**Definisi:** Mengambil data stok berdasarkan product ID.

**GET** `/stocks/product/:product_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| product_id | UUID (path) | ✅ | ID produk |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Stok tidak ditemukan |

---

### 8.5 Get Stock By Supplier ID
**Definisi:** Mengambil semua stok milik supplier tertentu.

**GET** `/stocks/supplier/:supplier_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| supplier_id | UUID (path) | ✅ | ID supplier |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |

---

### 8.6 Update Stock
**Definisi:** Memperbarui jumlah stok berdasarkan ID.

**PUT** `/stocks/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID stok |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| stock_quantity | int | ✅ | Jumlah stok baru (minimal 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Stok berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 8.7 Delete Stock
**Definisi:** Menghapus data stok berdasarkan ID (soft delete).

**DELETE** `/stocks/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID stok |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Stok berhasil dihapus |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `500` | Server error |

---

## 9. Orders

### 9.1 Create Order
**Definisi:** Membuat pesanan baru oleh SPPG dengan multiple item. Total harga dihitung otomatis.

**POST** `/orders`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| sppg_id | UUID | ✅ | ID SPPG |
| notes | string | ❌ | Catatan pesanan |
| items | array | ✅ | Minimal 1 item |
| items[].product_id | UUID | ✅ | ID produk |
| items[].quantity | int | ✅ | Jumlah (harus lebih dari 0) |

**Contoh Request:**
```json
{
    "sppg_id": "uuid-sppg",
    "notes": "Tolong kirim pagi hari",
    "items": [
        {
            "product_id": "uuid-produk-1",
            "quantity": 5
        },
        {
            "product_id": "uuid-produk-2",
            "quantity": 10
        }
    ]
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Pesanan berhasil dibuat |
| `400` | Validasi gagal atau produk tidak ditemukan |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 9.2 Get All Orders
**Definisi:** Mengambil semua pesanan dengan filter dan pagination.

**GET** `/orders`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter (Query):**
| Parameter | Type | Required | Default | Keterangan |
|-----------|------|----------|---------|------------|
| status | string | ❌ | - | Filter status: `pending`, `processing`, `shipped`, `completed`, `cancelled` |
| sppg_id | UUID | ❌ | - | Filter berdasarkan SPPG |
| start_date | string | ❌ | - | Filter tanggal mulai (YYYY-MM-DD) |
| end_date | string | ❌ | - | Filter tanggal akhir (YYYY-MM-DD) |
| page | int | ❌ | 1 | Halaman |
| limit | int | ❌ | 10 | Jumlah data per halaman |

**Contoh:**
```
GET /api/v1/orders?status=pending&page=1&limit=10
GET /api/v1/orders?start_date=2026-01-01&end_date=2026-03-31
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format parameter tidak valid |
| `401` | Token tidak valid |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Success",
    "data": [],
    "total": 50,
    "page": 1,
    "limit": 10
}
```

---

### 9.3 Get Order By ID
**Definisi:** Mengambil detail pesanan berdasarkan ID, termasuk detail item dan transaksi.

**GET** `/orders/:id`

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
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Pesanan tidak ditemukan |

---

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

---

### 9.5 Update Order Status
**Definisi:** Memperbarui status pesanan. Hanya pesanan dengan status `pending` atau `cancelled` yang bisa dihapus.

**PUT** `/orders/:id/status`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID pesanan |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_status | string | ✅ | `pending`, `processing`, `shipped`, `completed`, `cancelled` |

**Contoh Request:**
```json
{
    "order_status": "shipped"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

---

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

---

## 10. Transactions

### 10.1 Create Transaction
**Definisi:** Membuat transaksi pembayaran untuk pesanan tertentu.

**POST** `/transactions`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_id | UUID | ✅ | ID pesanan |
| payment_method | string | ❌ | Metode pembayaran (transfer, dll) |
| payment_proof | string | ❌ | URL bukti transfer |
| amount_paid | float64 | ✅ | Jumlah yang dibayar (harus lebih dari 0) |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `201` | Transaksi berhasil dibuat |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `500` | Server error |

---

### 10.2 Get All Transactions
**Definisi:** Mengambil semua data transaksi.

**GET** `/transactions`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |

---

### 10.3 Get Transaction By ID
**Definisi:** Mengambil data transaksi berdasarkan ID.

**GET** `/transactions/:id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID transaksi |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Transaksi tidak ditemukan |

---

### 10.4 Get Transaction By Order ID
**Definisi:** Mengambil data transaksi berdasarkan order ID.

**GET** `/transactions/order/:order_id`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| order_id | UUID (path) | ✅ | ID pesanan |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `400` | Format ID tidak valid |
| `401` | Token tidak valid |
| `404` | Transaksi tidak ditemukan |

---

### 10.5 Update Payment Status
**Definisi:** Memperbarui status pembayaran. Jika `paid`, order status otomatis berubah ke `processing`. Jika `failed`, order status berubah ke `cancelled`. Email notifikasi dikirim ke SPPG.

**PUT** `/transactions/:id/status`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID transaksi |

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| payment_status | string | ✅ | `unpaid`, `waiting_confirmation`, `paid`, `failed` |

**Contoh Request:**
```json
{
    "payment_status": "paid"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Status pembayaran berhasil diupdate |
| `400` | Validasi gagal |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

---

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

---

## 11. Payment (Midtrans)

### 11.1 Create Payment
**Definisi:** Membuat payment token Midtrans Snap untuk pesanan. Mengembalikan token dan redirect URL untuk menampilkan halaman pembayaran Midtrans.

**POST** `/payment/create`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | application/json | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `sppg`, `admin`

**Body:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| order_id | UUID | ✅ | ID pesanan |

**Contoh Request:**
```json
{
    "order_id": "uuid-pesanan"
}
```

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Token payment berhasil dibuat |
| `400` | Validasi gagal atau pesanan tidak ditemukan |
| `401` | Token tidak valid |
| `403` | Role tidak diizinkan |
| `500` | Gagal membuat payment di Midtrans |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Payment created successfully",
    "data": {
        "token": "248eb2a6-622e-4828-97a4-23266c0d1899",
        "redirect_url": "https://app.sandbox.midtrans.com/snap/v4/redirection/248eb2a6-622e-4828-97a4-23266c0d1899"
    }
}
```

---

## 12. Upload

### 12.1 Upload Supplier Document
**Definisi:** Upload dokumen verifikasi supplier (NIB, Halal, atau Other). Maksimal 5MB, format PDF/JPG/PNG.

**PATCH** `/upload/supplier/:id/document`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | multipart/form-data | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID supplier |

**Form Data:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| document_type | string | ✅ | `nib`, `halal`, atau `other` |
| file | file | ✅ | PDF, JPG, atau PNG, maks 5MB |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Dokumen berhasil diupload |
| `400` | Tipe file tidak valid, ukuran melebihi batas, atau document_type salah |
| `401` | Token tidak valid |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Document uploaded successfully",
    "data": {
        "document_type": "nib",
        "file_url": "/uploads/documents/supplier/uuid.pdf"
    }
}
```

---

### 12.2 Upload Product Image
**Definisi:** Upload foto produk. Maksimal 5MB, format JPG/PNG/WEBP.

**PATCH** `/upload/product/:id/image`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Content-Type | multipart/form-data | ✅ |
| Authorization | Bearer `<access_token>` | ✅ |

**Parameter:**
| Parameter | Type | Required | Keterangan |
|-----------|------|----------|------------|
| id | UUID (path) | ✅ | ID produk |

**Form Data:**
| Field | Type | Required | Keterangan |
|-------|------|----------|------------|
| file | file | ✅ | JPG, PNG, atau WEBP, maks 5MB |

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Foto berhasil diupload |
| `400` | Tipe file tidak valid atau ukuran melebihi batas |
| `401` | Token tidak valid |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Product image uploaded successfully",
    "data": {
        "file_url": "/uploads/images/products/uuid.jpg"
    }
}
```

---

## 13. Dashboard

### 13.1 Get Dashboard Summary
**Definisi:** Mengambil ringkasan data untuk dashboard admin.

**GET** `/dashboard/summary`

**Header:**
| Key | Value | Required |
|-----|-------|----------|
| Authorization | Bearer `<access_token>` | ✅ |

**Role yang diizinkan:** `admin`

**Parameter:** Tidak ada

**Status:**
| Status Code | Keterangan |
|-------------|------------|
| `200` | Data berhasil diambil |
| `401` | Token tidak valid |
| `403` | Bukan admin |
| `500` | Server error |

**Contoh Response (200):**
```json
{
    "success": true,
    "message": "Success",
    "data": {
        "total_supplier": 50,
        "supplier_pending": 10,
        "supplier_approved": 35,
        "supplier_rejected": 5,
        "total_sppg": 120,
        "total_order": 500,
        "order_pending": 20,
        "order_processing": 15,
        "order_shipped": 10,
        "order_completed": 450,
        "order_cancelled": 5
    }
}
```

---

## 📌 Catatan untuk Frontend

### Menyimpan Token
```javascript
localStorage.setItem('access_token', response.access_token)
localStorage.setItem('refresh_token', response.refresh_token)
```

### Menggunakan Token
```javascript
const response = await fetch('/api/v1/users', {
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('access_token')}`,
        'Content-Type': 'application/json'
    }
})
```

### Handle Token Expired (401)
```javascript
const refresh = await fetch('/api/v1/auth/refresh', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
        refresh_token: localStorage.getItem('refresh_token')
    })
})
const data = await refresh.json()
localStorage.setItem('access_token', data.access_token)
localStorage.setItem('refresh_token', data.refresh_token)
```

### Multi-Role Redirect Setelah Login
```javascript
const { role } = response.data
if (role === 'admin') router.push('/admin/dashboard')
else if (role === 'supplier') router.push('/supplier/dashboard')
else if (role === 'sppg') router.push('/sppg/dashboard')
```

### Upload File
```javascript
const formData = new FormData()
formData.append('file', file)
formData.append('document_type', 'nib')

const response = await fetch('/api/v1/upload/supplier/uuid/document', {
    method: 'PATCH',
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('access_token')}`
        // Jangan set Content-Type untuk multipart/form-data
    },
    body: formData
})
```

### Integrasi Midtrans Snap
```javascript
// Setelah dapat token dari /payment/create
window.snap.pay(token, {
    onSuccess: (result) => { console.log('Payment success', result) },
    onPending: (result) => { console.log('Payment pending', result) },
    onError: (result) => { console.log('Payment error', result) },
    onClose: () => { console.log('Payment popup closed') }
})
```
> Tambahkan script Midtrans Snap di HTML:
> ```html
> <script src="https://app.sandbox.midtrans.com/snap/snap.js" data-client-key="YOUR_CLIENT_KEY"></script>
> ```
