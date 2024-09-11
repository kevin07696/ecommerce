package services

import (
	"context"

	"github.com/kevin07696/ecommerce/domain"
	"github.com/kevin07696/ecommerce/domain/auth/models"
)

func (s *Service) processOTP(ctx context.Context, email models.Email, otp models.OTP, task string) error {
	foundOTP, err := s.cacher.Get(ctx, email.ToString())
	if err != nil {
		if err == domain.ErrNotFound {
			return domain.ErrUnauthorized
		}
		return domain.ErrInternalServer
	}

	otpWithHeader := addOTPHeader(otp, task)
	if otpWithHeader != foundOTP {
		return domain.ErrUnauthorized
	}

	s.cacher.Delete(ctx, email.ToString())

	return nil
}
