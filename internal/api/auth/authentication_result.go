package auth

import (
	"time"

	"github.com/danielmoisa/trip-planner/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
