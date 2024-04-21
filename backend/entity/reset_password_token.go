package entity

import "time"

type ResetPasswordToken struct {
	Email       string
	Token       string
	ExpiredAt   time.Time
	NewPassword string
}
