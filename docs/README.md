# API Documentation - Pangantara

**Base URL:** `https://be-internship.bccdev.id/hanif/api/v1`
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

## 🔓 Public Routes

### Auth
- [Register](./api/auth/register.md)
- [Login](./api/auth/login.md)
- [Forgot Password](./api/auth/forgot-password.md)
- [Reset Password](./api/auth/reset-password.md)
- [Refresh Token](./api/auth/refresh-token.md)

### Webhook
- [Midtrans Webhook](./api/webhook/midtrans.md)

---

## 🔐 Protected Routes

### Dashboard
- [Get Dashboard Summary](./api/dashboard/get-dashboard-summary.md)

### Users
- [Get All Users](./api/users/get-all-user.md)
- [Get User By ID](./api/users/get-user-by-id.md)
- [Create User](./api/users/create-user.md)
- [Update User](./api/users/update-user.md)
- [Delete User](./api/users/delete-user.md)

### Suppliers
- [Get All Suppliers](./api/suppliers/get-all-supplier.md)
- [Get Supplier By ID](./api/suppliers/get-supplier-by-id.md)
- [Get Supplier By User ID](./api/suppliers/get-supplier-by-user-id.md)
- [Create Supplier](./api/suppliers/create-supplier.md)
- [Update Supplier](./api/suppliers/update-supplier.md)
- [Delete Supplier](./api/suppliers/delete-supplier.md)
- [Verify Supplier (Admin Only)](./api/suppliers/verify-supplier.md)

### Supplier Drafts
- [Get Draft](./api/supplier-drafts/get-draft.md)
- [Save Draft](./api/supplier-drafts/save-draft.md)
- [Submit Draft](./api/supplier-drafts/submit-draft.md)
- [Delete Draft](./api/supplier-drafts/delete-draft.md)
- [Upload Draft Document](./api/supplier-drafts/upload-draft-document.md)

### SPPG
- [Get All SPPG](./api/sppg/get-all-sppg.md)
- [Get SPPG By ID](./api/sppg/get-sppg-by-id.md)
- [Get SPPG By User ID](./api/sppg/get-sppg-by-user-id.md)
- [Create SPPG](./api/sppg/create-sppg.md)
- [Update SPPG](./api/sppg/update-sppg.md)
- [Delete SPPG](./api/sppg/delete-sppg.md)

### Products
- [Get All Products](./api/products/get-all-products.md)
- [Get Product By ID](./api/products/get-product-by-id.md)
- [Get Product By Supplier](./api/products/get-product-by-supplier.md)
- [Create Product](./api/products/create-product.md)
- [Update Product](./api/products/update-product.md)
- [Delete Product](./api/products/delete-product.md)

### Stock
- [Get All Stocks](./api/stock/get-all-stocks.md)
- [Get Stock By ID](./api/stock/get-stock-by-id.md)
- [Get Stock By Product ID](./api/stock/get-stock-by-product-id.md)
- [Get Stock By Supplier ID](./api/stock/get-stock-by-supplier-id.md)
- [Create Stock](./api/stock/create-stock.md)
- [Update Stock](./api/stock/update-stock.md)
- [Delete Stock](./api/stock/delete-stock.md)

### Orders
- [Get All Orders](./api/orders/get-all-orders.md)
- [Get Order By ID](./api/orders/get-order-by-id.md)
- [Get Order By SPPG ID](./api/orders/get-order-by-sppg-id.md)
- [Create Order](./api/orders/create-order.md)
- [Update Order Status](./api/orders/update-order-status.md)
- [Delete Order](./api/orders/delete-order.md)

### Transactions
- [Get All Transactions](./api/transactions/get-all-transactions.md)
- [Get Transaction By ID](./api/transactions/get-transaction-by-id.md)
- [Get Transaction By Order ID](./api/transactions/get-transaction-by-order-id.md)
- [Create Transaction](./api/transactions/create-transaction.md)
- [Update Payment Status](./api/transactions/update-payment-status.md)
- [Delete Transaction](./api/transactions/delete-transaction.md)

### Payments
- [Create Payment](./api/payments/create-payment.md)

### Uploads
- [Upload Product Image](./api/uploads/upload-product-image.md)
- [Upload Supplier Document](./api/uploads/upload-supplier-document.md)