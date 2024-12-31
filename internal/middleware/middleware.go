package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-service-template/config"
	"go-service-template/internal/infrastructure/adapter/cas"
	"go-service-template/internal/infrastructure/adapter/cas/casDTO"
	errors "go-service-template/pkgs/utils/common/servErr"
)

type AuthMiddleware struct {
	cfg *config.AppConfig
	cas cas.CasAdapter
}

func NewAuthMiddleware(cfg *config.AppConfig, cas cas.CasAdapter) *AuthMiddleware {
	return &AuthMiddleware{
		cfg: cfg,
		cas: cas,
	}
}

func (auth *AuthMiddleware) RequiredAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//read token from header
		token := c.Get("Authorization")
		if token == "" {
			return errors.ErrUnauthenticated
		}

		bearerToken := token[7:]
		userAuth, err := auth.cas.VerifyToken(c.Context(), &casDTO.CASVerifyTokenReq{
			Token: bearerToken,
		})
		if err != nil {
			return errors.ErrUnauthenticated
		}

		//set user auth to context
		c.Locals(UserIdAuthContext, userAuth.UserId)
		c.Locals(UsernameAuthContext, userAuth.Username)
		c.Locals(RolesAuthContext, userAuth.Roles)

		return c.Next()
	}
}
