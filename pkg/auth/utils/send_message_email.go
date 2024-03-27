package utils

import (
	"fmt"
	"net/smtp"
	"social_network_for_programmers/internal/entity/emails"
)

func SendMessageEmail(s *emails.MessageEmail) error {
	ath := smtp.PlainAuth("", s.Ath.Username, s.Ath.Password, s.Ath.Host)

	if err := smtp.SendMail(s.Ath.Host+":"+s.Ath.Port, ath, s.From, s.To, s.Content); err != nil {
		return fmt.Errorf("failed to send a message: %s", err.Error())
	}

	return nil
}
