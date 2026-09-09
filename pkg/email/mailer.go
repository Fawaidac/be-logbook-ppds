package email

import (
	"fmt"
	"log"
	"net/smtp"
	"be-logbook-ppds/configs"
)
	
type Mailer interface {
	SendApprovalEmail(toEmail, recipientName, username, defaultPassword string) error
	SendRejectionEmail(toEmail, recipientName, reason string) error
}

type mailer struct {
	cfg *configs.Config
}

func NewMailer(cfg *configs.Config) Mailer {
	return &mailer{cfg: cfg}
}

func (m *mailer) sendMail(toEmail, subject, bodyHTML string) error {
	if m.cfg.SMTPHost == "" || m.cfg.SMTPUser == "" {
		log.Printf("[EMAIL NOTIFICATION LOG] SMTP Not Configured. Email simulated for %s | Subject: %s", toEmail, subject)
		return nil
	}

	from := m.cfg.SMTPFrom
	if from == "" {
		from = m.cfg.SMTPUser
	}

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + bodyHTML

	auth := smtp.PlainAuth("", m.cfg.SMTPUser, m.cfg.SMTPPassword, m.cfg.SMTPHost)
	addr := fmt.Sprintf("%s:%s", m.cfg.SMTPHost, m.cfg.SMTPPort)

	err := smtp.SendMail(addr, auth, from, []string{toEmail}, []byte(message))
	if err != nil {
		log.Printf("[EMAIL ERROR] Failed sending email to %s: %v", toEmail, err)
		return err
	}

	log.Printf("[EMAIL SUCCESS] Email sent to %s", toEmail)
	return nil
}

func (m *mailer) SendApprovalEmail(toEmail, recipientName, username, defaultPassword string) error {
	subject := "Persetujuan Akun PPDS - Logbook RSUD dr. Soebandi"
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head><meta charset="utf-8"></head>
	<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
			<h2 style="color: #2e7d32; border-bottom: 2px solid #2e7d32; padding-bottom: 10px;">Pendaftaran Akun PPDS Disetujui ✓</h2>
			<p>Yth. <strong>%s</strong>,</p>
			<p>Selamat! Permintaan pendaftaran akun PPDS Anda untuk sistem Logbook RSUD dr. Soebandi telah <strong>disetujui oleh Admin</strong>.</p>
			<div style="background-color: #f1f8e9; padding: 15px; border-left: 4px solid #2e7d32; margin: 20px 0;">
				<h4 style="margin-top: 0; color: #2e7d32;">Informasi Akun Anda:</h4>
				<p style="margin: 5px 0;"><strong>Username:</strong> <code>%s</code></p>
				<p style="margin: 5px 0;"><strong>Password:</strong> Gunakan kata sandi yang Anda daftarkan saat registrasi.</p>
			</div>
			<p>Silakan masuk ke platform Logbook PPDS menggunakan username dan kata sandi yang sudah Anda buat saat mendaftar.</p>
			<hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
			<p style="font-size: 12px; color: #777;">Email ini dikirimkan secara otomatis oleh Sistem Logbook PPDS RSUD dr. Soebandi.</p>
		</div>
	</body>
	</html>
	`, recipientName, username)

	return m.sendMail(toEmail, subject, body)
}

func (m *mailer) SendRejectionEmail(toEmail, recipientName, reason string) error {
	subject := "Pemberitahuan Status Pendaftaran PPDS - Logbook RSUD dr. Soebandi"
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head><meta charset="utf-8"></head>
	<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
		<div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px;">
			<h2 style="color: #c62828; border-bottom: 2px solid #c62828; padding-bottom: 10px;">Pendaftaran Akun PPDS Belum Disetujui</h2>
			<p>Yth. <strong>%s</strong>,</p>
			<p>Mohon maaf, permintaan pendaftaran akun PPDS Anda untuk sistem Logbook RSUD dr. Soebandi <strong>belum dapat disetujui</strong> oleh Admin.</p>
			<div style="background-color: #ffebee; padding: 15px; border-left: 4px solid #c62828; margin: 20px 0;">
				<h4 style="margin-top: 0; color: #c62828;">Catatan / Alasan Penolakan:</h4>
				<p style="margin: 5px 0; font-style: italic;">"%s"</p>
			</div>
			<p>Silakan lakukan pendaftaran ulang dengan melengkapi atau memperbaiki dokumen/data sesuai dengan alasan di atas.</p>
			<hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
			<p style="font-size: 12px; color: #777;">Email ini dikirimkan secara otomatis oleh Sistem Logbook PPDS RSUD dr. Soebandi.</p>
		</div>
	</body>
	</html>
	`, recipientName, reason)

	return m.sendMail(toEmail, subject, body)
}
