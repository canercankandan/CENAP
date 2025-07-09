package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/gomail.v2"
	"cenap/internal/models"
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

// SendSupportChatNotification, yeni ziyaretçi canlı desteğe geldiğinde mail bildirimi gönderir
func (es *EmailService) SendSupportChatNotification(adminEmail, visitorName, sessionID, userAgent string) error {
	if es.dialer == nil {
		log.Printf("E-posta gönderimi devre dışı. Support chat bildirimi: %s - %s", visitorName, sessionID)
		return nil
	}

	subject := "Yeni Canlı Destek Ziyaretçisi - Cenap Su Arıtma"
	body := fmt.Sprintf(`
		<h2>🔔 Yeni Canlı Destek Ziyaretçisi</h2>
		<p>Merhaba,</p>
		<p><strong>%s</strong> adlı ziyaretçi canlı destek sayfasına girdi ve sizi bekliyor.</p>
		<br>
		<p><strong>Ziyaretçi Bilgileri:</strong></p>
		<ul>
			<li><strong>Ad:</strong> %s</li>
			<li><strong>Session ID:</strong> %s</li>
			<li><strong>Tarih:</strong> %s</li>
			<li><strong>Tarayıcı:</strong> %s</li>
		</ul>
		<br>
		<p>Ziyaretçiye yardımcı olmak için admin panelini kontrol edin.</p>
		
		<div style="text-align: center; margin: 20px 0;">
			<a href="https://xn--suartmauzman-44bi.com/admin" style="display: inline-block; background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold; margin: 10px;">📱 Admin Paneli</a>
			<a href="https://xn--suartmauzman-44bi.com/admin/support" style="display: inline-block; background-color: #28a745; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold; margin: 10px;">💬 Canlı Destek</a>
		</div>
		
		<p><strong>Hızlı Erişim Linkleri:</strong></p>
		<ul>
			<li><a href="https://xn--suartmauzman-44bi.com/admin">Ana Admin Paneli</a></li>
			<li><a href="https://xn--suartmauzman-44bi.com/admin/support">Canlı Destek Paneli</a></li>
		</ul>
		
		<br>
		<p><em>Bu bildirim, ziyaretçi canlı destek sayfasına girdiğinde otomatik olarak gönderilmiştir.</em></p>
		<br>
		<p>Saygılarımızla,<br>Cenap Su Arıtma</p>
	`, visitorName, visitorName, sessionID, time.Now().Format("02.01.2006 15:04:05"), userAgent)

	m := gomail.NewMessage()
	m.SetHeader("From", es.from)
	m.SetHeader("To", adminEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return es.dialer.DialAndSend(m)
}

// SendAdminOrderNotification, yeni sipariş geldiğinde admin'e bildirim gönderir
func (es *EmailService) SendAdminOrderNotification(adminEmail string, order *models.Order) error {
	if es.dialer == nil {
		log.Printf("E-posta gönderimi devre dışı. Admin sipariş bildirimi: %s", order.OrderNumber)
		return nil
	}

	// Sipariş ürünlerini formatla
	var itemsHTML string
	for _, item := range order.Items {
		itemsHTML += fmt.Sprintf(`
			<tr>
				<td style="padding: 10px; border-bottom: 1px solid #eee;">%s</td>
				<td style="padding: 10px; border-bottom: 1px solid #eee;">%.2f TL</td>
				<td style="padding: 10px; border-bottom: 1px solid #eee;">%d</td>
				<td style="padding: 10px; border-bottom: 1px solid #eee;">%.2f TL</td>
			</tr>
		`, item.Name, item.Price, item.Quantity, item.Price*float64(item.Quantity))
	}

	subject := "Yeni Sipariş Bildirimi - Cenap Su Arıtma"
	body := fmt.Sprintf(`
		<h2>🛒 Yeni Sipariş Alındı!</h2>
		<p>Merhaba,</p>
		<p>Yeni bir sipariş alındı ve işleme hazır.</p>
		<br>
		<p><strong>Sipariş Detayları:</strong></p>
		<ul>
			<li><strong>Sipariş Numarası:</strong> %s</li>
			<li><strong>Müşteri Adı:</strong> %s</li>
			<li><strong>E-posta:</strong> %s</li>
			<li><strong>Telefon:</strong> %s</li>
			<li><strong>Adres:</strong> %s</li>
			<li><strong>Ödeme Yöntemi:</strong> %s</li>
			<li><strong>Toplam Tutar:</strong> %.2f TL</li>
			<li><strong>Tarih:</strong> %s</li>
		</ul>
		<br>
		<p><strong>Sipariş Ürünleri:</strong></p>
		<table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
			<thead>
				<tr style="background-color: #f8f9fa;">
					<th style="padding: 10px; text-align: left; border-bottom: 2px solid #dee2e6;">Ürün</th>
					<th style="padding: 10px; text-align: left; border-bottom: 2px solid #dee2e6;">Fiyat</th>
					<th style="padding: 10px; text-align: left; border-bottom: 2px solid #dee2e6;">Adet</th>
					<th style="padding: 10px; text-align: left; border-bottom: 2px solid #dee2e6;">Toplam</th>
				</tr>
			</thead>
			<tbody>
				%s
			</tbody>
		</table>
		<br>
		<p><strong>Müşteri Notları:</strong></p>
		<p style="background-color: #f8f9fa; padding: 15px; border-radius: 5px; border-left: 4px solid #007bff;">%s</p>
		<br>
		<div style="text-align: center; margin: 20px 0;">
			<a href="https://xn--suartmauzman-44bi.com/admin" style="display: inline-block; background-color: #007bff; color: white; padding: 12px 24px; text-decoration: none; border-radius: 5px; font-weight: bold; margin: 10px;">📱 Admin Paneli</a>
		</div>
		<br>
		<p><strong>Hızlı Erişim Linkleri:</strong></p>
		<ul>
			<li><a href="https://xn--suartmauzman-44bi.com/admin">Ana Admin Paneli</a></li>
		</ul>
		<br>
		<p><em>Bu bildirim, müşteri siparişi tamamladığında otomatik olarak gönderilmiştir.</em></p>
		<br>
		<p>Saygılarımızla,<br>Cenap Su Arıtma</p>
	`, order.OrderNumber, order.CustomerName, order.Email, order.Phone, order.Address, order.PaymentMethod, order.TotalPrice, order.CreatedAt.Format("02.01.2006 15:04:05"), itemsHTML, order.Notes)

	m := gomail.NewMessage()
	m.SetHeader("From", es.from)
	m.SetHeader("To", adminEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return es.dialer.DialAndSend(m)
} 