package services

import (
	"context"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
)

func (s *Service) DeleteSession(ctx context.Context, w http.ResponseWriter, r *http.Request) *domain.CustomError {
	err := s.sessionManager.Delete(ctx, w, r)
	if err != nil {
		return domain.CustomizeError(domain.ErrInternalServer, ErrMsgInternalServer)
	}

	return nil
}
