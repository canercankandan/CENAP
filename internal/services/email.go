package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/gomail.v2"
)

// EmailService, e-posta gönderimi için kullanılır
type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

// NewEmailService, yeni bir EmailService örneği oluşturur
func NewEmailService() *EmailService {
	// Gmail SMTP ayarları
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUser := "irmaksuaritmam@gmail.com" // Gmail adresiniz
	smtpPass := "smve btgb zoih rkkd"      // Gmail uygulama şifresi

	// Eğer environment variable'lar ayarlanmışsa, onları kullan
	if envUser := os.Getenv("SMTP_USER"); envUser != "" {
		smtpUser = envUser
	}
	if envPass := os.Getenv("SMTP_PASS"); envPass != "" {
		smtpPass = envPass
	}

	// SMTP bilgileri kontrol et
	if smtpUser == "" || smtpPass == "" {
		log.Println("SMTP bilgileri ayarlanmamış. E-posta gönderimi devre dışı.")
		return &EmailService{
			dialer: nil,
			from:   "noreply@cenap.com",
		}
	}

	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	return &EmailService{
		dialer: dialer,
		from:   smtpUser,
	}
}

// SendPasswordResetEmail, şifre sıfırlama e-postası gönderir
func (es *EmailService) SendPasswordResetEmail(to, token string) error {
	if es.dialer == nil {
		// SMTP ayarlanmamışsa, sadece log'a yaz
		log.Printf("E-posta gönderimi devre dışı. Şifre sıfırlama token'ı: %s", token)
		return nil
	}

	subject := "Şifre Sıfırlama - Cenap Su Arıtma"
	body := fmt.Sprintf(`
		<h2>Şifre Sıfırlama İsteği</h2>
		<p>Merhaba,</p>
		<p>Şifrenizi sıfırlamak için aşağıdaki bağlantılardan birini kullanın:</p>
		
		<div style="text-align: center; margin: 20px 0;">
			<a href="https://xn--suartmauzman-44bi.com/reset-password?token=%s" style="display: inline-block; background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold; margin: 10px;">🔐 Şifremi Sıfırla</a>
		</div>
		
		<p><strong>Alternatif Link:</strong></p>
		<p><a href="https://xn--suartmauzman-44bi.com/reset-password?token=%s">https://xn--suartmauzman-44bi.com/reset-password?token=%s</a></p>
		
		<p>Bu bağlantı 1 saat süreyle geçerlidir.</p>
		<p>Eğer bu isteği siz yapmadıysanız, bu e-postayı görmezden gelebilirsiniz.</p>
		<br>
		<p>Saygılarımızla,<br>Cenap Su Arıtma</p>
	`, token, token, token)

	m := gomail.NewMessage()
	m.SetHeader("From", es.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return es.dialer.DialAndSend(m)
}

// SendWelcomeEmail, hoş geldin e-postası gönderir
func (es *EmailService) SendWelcomeEmail(to, username string) error {
	if es.dialer == nil {
		log.Printf("E-posta gönderimi devre dışı. Hoş geldin e-postası: %s", username)
		return nil
	}

	subject := "Hoş Geldiniz - Cenap Su Arıtma"
	body := fmt.Sprintf(`
		<h2>Hoş Geldiniz!</h2>
		<p>Merhaba <strong>%s</strong>,</p>
		<p>Cenap Su Arıtma ailesine hoş geldiniz! Hesabınız başarıyla oluşturuldu.</p>
		<br>
		<p>Artık aşağıdaki hizmetlerimizden yararlanabilirsiniz:</p>
		<ul>
			<li>✅ Kaliteli su arıtma cihazları</li>
			<li>✅ 7/24 canlı destek</li>
			<li>✅ Ücretsiz kurulum</li>
			<li>✅ 5 yıl garanti</li>
		</ul>
		<br>
		<div style="text-align: center; margin: 20px 0;">
			<a href="https://xn--suartmauzman-44bi.com" style="display: inline-block; background-color: #28a745; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold;">🏠 Ana Sayfaya Git</a>
		</div>
		<br>
		<p>Herhangi bir sorunuz olursa <a href="https://xn--suartmauzman-44bi.com/contact">iletişim</a> sayfamızdan bize ulaşabilirsiniz.</p>
		<br>
		<p>Saygılarımızla,<br>Cenap Su Arıtma</p>
	`, username)

	m := gomail.NewMessage()
	m.SetHeader("From", es.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return es.dialer.DialAndSend(m)
}

// SendVideoCallNotification, video call talebi bildirimi gönderir
func (es *EmailService) SendVideoCallNotification(adminEmail, customerName, sessionID string) error {
	if es.dialer == nil {
		log.Printf("E-posta gönderimi devre dışı. Video call bildirimi: %s - %s", customerName, sessionID)
		return nil
	}

	subject := "Video Görüşme Talebi - Cenap Su Arıtma"
	body := fmt.Sprintf(`
		<h2>Video Görüşme Talebi</h2>
		<p>Merhaba,</p>
		<p><strong>%s</strong> adlı müşteri canlı destek üzerinden video görüşme talebinde bulundu.</p>
		<p><strong>Session ID:</strong> %s</p>
		<p><strong>Tarih:</strong> %s</p>
		<br>
		<p>Video görüşme talebini yanıtlamak için admin panelini kontrol edin.</p>
		
		<div style="text-align: center; margin: 20px 0;">
			<a href="https://xn--suartmauzman-44bi.com/admin/support" style="display: inline-block; background-color: #28a745; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold; margin: 10px;">📱 Admin Destek Paneli</a>
		</div>
		
		<p><strong>Alternatif Linkler:</strong></p>
		<ul>
			<li><a href="https://xn--suartmauzman-44bi.com/admin">Ana Admin Paneli</a></li>
			<li><a href="https://xn--suartmauzman-44bi.com/admin/support">Destek Paneli</a></li>
		</ul>
		
		<br>
		<p>Saygılarımızla,<br>Cenap Su Arıtma</p>
	`, customerName, sessionID, time.Now().Format("02.01.2006 15:04:05"))

	m := gomail.NewMessage()
	m.SetHeader("From", es.from)
	m.SetHeader("To", adminEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return es.dialer.DialAndSend(m)
} 