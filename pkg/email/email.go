package email

import (
	"fmt"
	"sppg-backend/config"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	cfg := config.AppConfig

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", cfg.SMTPName, cfg.SMTPEmail))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPEmail, cfg.SMTPPassword)

	return d.DialAndSend(m)
}

func SendForgotPasswordEmail(to, name, resetLink string) error {
	subject := "Reset Password - Pangantara"
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
			<div style="background-color: #2D6A4F; padding: 20px; text-align: center;">
				<h1 style="color: white; margin: 0;">PANGANTARA</h1>
			</div>
			<div style="padding: 30px; background-color: #f9f9f9;">
				<h2>Halo, %s!</h2>
				<p>Kami menerima permintaan untuk mereset kata sandi akun Pangantara kamu.</p>
				<p>Klik tombol di bawah untuk mereset kata sandi kamu:</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="%s" style="background-color: #F4A261; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
						Reset Kata Sandi
					</a>
				</div>
				<p>Link ini akan kadaluarsa dalam <strong>15 menit</strong>.</p>
				<p>Jika kamu tidak meminta reset kata sandi, abaikan email ini.</p>
			</div>
			<div style="padding: 20px; text-align: center; color: #888;">
				<p>© 2026 Pangantara. All rights reserved.</p>
			</div>
		</div>
	`, name, resetLink)

	return SendEmail(to, subject, body)
}

func SendSupplierApprovedEmail(to, name, storeName string) error {
	subject := "Selamat! Pendaftaran Supplier Disetujui - Pangantara"
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
			<div style="background-color: #2D6A4F; padding: 20px; text-align: center;">
				<h1 style="color: white; margin: 0;">PANGANTARA</h1>
			</div>
			<div style="padding: 30px; background-color: #f9f9f9;">
				<h2>Selamat, %s! 🎉</h2>
				<p>Pendaftaran supplier <strong>%s</strong> kamu telah <strong style="color: #2D6A4F;">DISETUJUI</strong> oleh admin Pangantara.</p>
				<p>Kamu sekarang dapat:</p>
				<ul>
					<li>Menambahkan produk ke platform</li>
					<li>Mengelola stok produk</li>
					<li>Menerima pesanan dari SPPG</li>
				</ul>
				<div style="text-align: center; margin: 30px 0;">
					<a href="http://localhost:3000/supplier/dashboard" style="background-color: #2D6A4F; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
						Masuk ke Dashboard
					</a>
				</div>
			</div>
			<div style="padding: 20px; text-align: center; color: #888;">
				<p>© 2026 Pangantara. All rights reserved.</p>
			</div>
		</div>
	`, name, storeName)

	return SendEmail(to, subject, body)
}

func SendSupplierRejectedEmail(to, name, storeName string, notes *string) error {
	subject := "Pendaftaran Supplier Ditolak - Pangantara"

	reason := "Tidak ada catatan tambahan dari admin."
	if notes != nil && *notes != "" {
		reason = *notes
	}

	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
			<div style="background-color: #2D6A4F; padding: 20px; text-align: center;">
				<h1 style="color: white; margin: 0;">PANGANTARA</h1>
			</div>
			<div style="padding: 30px; background-color: #f9f9f9;">
				<h2>Halo, %s</h2>
				<p>Mohon maaf, pendaftaran supplier <strong>%s</strong> kamu <strong style="color: #e63946;">DITOLAK</strong> oleh admin Pangantara.</p>
				<div style="background-color: #fff3f3; border-left: 4px solid #e63946; padding: 15px; margin: 20px 0;">
					<p style="margin: 0;"><strong>Alasan penolakan:</strong></p>
					<p style="margin: 5px 0 0 0;">%s</p>
				</div>
				<p>Kamu dapat mendaftar ulang setelah memperbaiki kekurangan di atas.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="http://localhost:3000/register/supplier" style="background-color: #F4A261; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
						Daftar Ulang
					</a>
				</div>
			</div>
			<div style="padding: 20px; text-align: center; color: #888;">
				<p>© 2026 Pangantara. All rights reserved.</p>
			</div>
		</div>
	`, name, storeName, reason)

	return SendEmail(to, subject, body)
}

func SendPaymentConfirmedEmail(to, name, orderID string, totalAmount float64) error {
	subject := "Pembayaran Dikonfirmasi - Pangantara"
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
			<div style="background-color: #2D6A4F; padding: 20px; text-align: center;">
				<h1 style="color: white; margin: 0;">PANGANTARA</h1>
			</div>
			<div style="padding: 30px; background-color: #f9f9f9;">
				<h2>Halo, %s! 🎉</h2>
				<p>Pembayaran kamu telah <strong style="color: #2D6A4F;">DIKONFIRMASI</strong> oleh admin Pangantara.</p>
				<div style="background-color: #fff; border: 1px solid #e0e0e0; border-radius: 8px; padding: 20px; margin: 20px 0;">
					<p><strong>Detail Pembayaran:</strong></p>
					<p>Order ID: <strong>%s</strong></p>
					<p>Total Pembayaran: <strong>Rp %s</strong></p>
					<p>Status: <strong style="color: #2D6A4F;">LUNAS</strong></p>
				</div>
				<p>Pesanan kamu sedang diproses dan akan segera dikirim.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="http://localhost:3000/sppg/orders" style="background-color: #2D6A4F; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
						Lihat Pesanan
					</a>
				</div>
			</div>
			<div style="padding: 20px; text-align: center; color: #888;">
				<p>© 2026 Pangantara. All rights reserved.</p>
			</div>
		</div>
	`, name, orderID, formatRupiah(totalAmount))

	return SendEmail(to, subject, body)
}

func SendPaymentRejectedEmail(to, name, orderID string, totalAmount float64) error {
	subject := "Pembayaran Ditolak - Pangantara"
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
			<div style="background-color: #2D6A4F; padding: 20px; text-align: center;">
				<h1 style="color: white; margin: 0;">PANGANTARA</h1>
			</div>
			<div style="padding: 30px; background-color: #f9f9f9;">
				<h2>Halo, %s</h2>
				<p>Mohon maaf, pembayaran kamu <strong style="color: #e63946;">DITOLAK</strong> oleh admin Pangantara.</p>
				<div style="background-color: #fff3f3; border: 1px solid #e63946; border-radius: 8px; padding: 20px; margin: 20px 0;">
					<p><strong>Detail Pembayaran:</strong></p>
					<p>Order ID: <strong>%s</strong></p>
					<p>Total Pembayaran: <strong>Rp %s</strong></p>
					<p>Status: <strong style="color: #e63946;">DITOLAK</strong></p>
				</div>
				<p>Kemungkinan penyebab penolakan:</p>
				<ul>
					<li>Bukti transfer tidak jelas</li>
					<li>Jumlah transfer tidak sesuai</li>
					<li>Transfer ke rekening yang salah</li>
				</ul>
				<p>Silakan upload ulang bukti transfer yang benar.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="http://localhost:3000/sppg/orders" style="background-color: #F4A261; color: white; padding: 12px 30px; text-decoration: none; border-radius: 5px; font-weight: bold;">
						Upload Ulang Bukti Transfer
					</a>
				</div>
			</div>
			<div style="padding: 20px; text-align: center; color: #888;">
				<p>© 2026 Pangantara. All rights reserved.</p>
			</div>
		</div>
	`, name, orderID, formatRupiah(totalAmount))

	return SendEmail(to, subject, body)
}

func formatRupiah(amount float64) string {
	return fmt.Sprintf("%.0f", amount)
}