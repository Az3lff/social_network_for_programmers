package authentication

import "social_network_for_programmers/internal/service/authentication/models"

type AuthStorage struct {
	Users map[int]models.User
}
