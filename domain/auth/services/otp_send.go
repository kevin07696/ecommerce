package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

type SendOTPReq struct {
	Email string
	Key   string
	Task  string
}

func (s *Service) sendOTP(ctx context.Context, email models.Email, task string) error {
	otp := generateOTP()

	headerOTP := addOTPHeader(models.OTP(otp), task)
	err := s.cacher.Set(ctx, email.ToString(), headerOTP, time.Hour)
	if err != nil {
		return domain.ErrInternalServer
	}

	subject := "Verify Your Email Address for Your Account"
	body := fmt.Sprintf("Please use the one time password below:\n\n%s\n\nThank you,\nThe Team", otp)

	err = s.emailer.SendEmail(ctx, email.ToString(), subject, body)
	if err != nil {
		return domain.ErrInternalServer
	}

	log.Println("[Success] OTP: sent OTP to email and cache")

	return nil
}

func addOTPHeader(otp models.OTP, task string) string {
	return task + string(otp)
}

func generateOTP() string {
	// length of bytes = desired characters * 3/4
	const length = 6
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
