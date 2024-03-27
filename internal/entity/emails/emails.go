package emails

import "social_network_for_programmers/internal/config"

type MessageEmail struct {
	Ath     *config.AuthEmail
	From    string
	To      []string
	Content []byte
}
