package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type EmailTemplate struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailTemplates struct {
	Verification  EmailTemplate `json:"verification"`
	PasswordReset EmailTemplate `json:"passwordReset"`
}

type SMTPConfig struct {
	Host          string
	Port          int
	Username      string
	Password      string
	From          string
	TemplatesPath string
}

type SMTPEmailSender struct {
	config    SMTPConfig
	templates EmailTemplates
}

func NewSMTPEmailSender(config SMTPConfig) (*SMTPEmailSender, error) {
	sender := &SMTPEmailSender{config: config}

	if config.TemplatesPath != "" {
		if err := sender.loadTemplates(config.TemplatesPath); err != nil {
			return nil, fmt.Errorf("failed to load email templates: %w", err)
		}
	} else {
		sender.templates = defaultTemplates()
	}

	return sender, nil
}

func (s *SMTPEmailSender) loadTemplates(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.templates)
}

func defaultTemplates() EmailTemplates {
	return EmailTemplates{
		Verification: EmailTemplate{
			Subject: "이메일 인증 코드",
			Body:    "인증 코드: {{.Code}}",
		},
		PasswordReset: EmailTemplate{
			Subject: "비밀번호 재설정 코드",
			Body:    "인증 코드: {{.Code}}",
		},
	}
}

func (s *SMTPEmailSender) renderTemplate(tmplStr string, data map[string]string) (string, error) {
	tmpl, err := template.New("email").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (s *SMTPEmailSender) SendVerificationEmail(email, code string) error {
	body, err := s.renderTemplate(s.templates.Verification.Body, map[string]string{"Code": code})
	if err != nil {
		return err
	}
	return s.sendEmail(email, s.templates.Verification.Subject, body)
}

func (s *SMTPEmailSender) SendPasswordResetEmail(email, code string) error {
	body, err := s.renderTemplate(s.templates.PasswordReset.Body, map[string]string{"Code": code})
	if err != nil {
		return err
	}
	return s.sendEmail(email, s.templates.PasswordReset.Subject, body)
}

func (s *SMTPEmailSender) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		to, subject, body))

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	return smtp.SendMail(addr, auth, s.config.From, []string{to}, msg)
}

// NoopEmailSender 개발/테스트용 이메일 발송기 (실제로 발송하지 않음)
type NoopEmailSender struct{}

func NewNoopEmailSender() *NoopEmailSender {
	return &NoopEmailSender{}
}

func (n *NoopEmailSender) SendVerificationEmail(email, code string) error {
	fmt.Printf("[DEV] Verification email to %s: code=%s\n", email, code)
	return nil
}

func (n *NoopEmailSender) SendPasswordResetEmail(email, code string) error {
	fmt.Printf("[DEV] Password reset email to %s: code=%s\n", email, code)
	return nil
}
