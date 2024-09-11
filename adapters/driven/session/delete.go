package session

import (
	"context"
	"log"
	"net/http"

	"github.com/kevin07696/ecommerce/domain"
)

func (s SessionManager) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := s.manager.Destroy(ctx, w, r)
	if err != nil {
		log.Printf("Fails to delete session: %v, %v", domain.ErrInternalServer, err)
		return domain.ErrInternalServer
	}
	return nil
}
