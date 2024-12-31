package cas

import (
	"context"
	"go-service-template/internal/infrastructure/adapter/cas/casDTO"
)

type CasAdapter interface {
	VerifyToken(ctx context.Context, req *casDTO.CASVerifyTokenReq) (*casDTO.CASVerifyTokenResp, error)
}
