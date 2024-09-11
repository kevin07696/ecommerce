package models

import (
	"strings"
)

type OTPReq string

type OTP string

func (m Models) NewOTP(otp string) (OTP, error) {
	otp = strings.TrimSpace(otp)
	if otp == "" {
		return "", ErrEmptyOTP
	}
	if len(otp) != 8 {
		return "", ErrInvalidOTP
	}
	return OTP(otp), nil
}
