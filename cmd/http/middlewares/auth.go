package middlewares

import (
	"errors"
	"go-skeleton/pkg/jwtExtractor"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/roles"
	"go-skeleton/pkg/signerVerifier"

	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	logger    *logger.Logger
	signer    *signerVerifier.Signer
	acl       *roles.CasbinRule
	extractor *jwtExtractor.JWTExtractor
}

func NewAuthMiddleware(reg *registry.Registry) *AuthMiddleware {
	return &AuthMiddleware{
		logger:    reg.Inject("logger").(*logger.Logger),
		signer:    reg.Inject("signerVerifier").(*signerVerifier.Signer),
		acl:       reg.Inject("roles").(*roles.CasbinRule),
		extractor: reg.Inject("jwtExtractor").(*jwtExtractor.JWTExtractor),
	}
}

func (m *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Authorization := c.Request().Header.Get("Authorization")
		if Authorization == "" {
			return c.JSON(401, "Not Authorized")
		}

		verify, verifyErr := m.signer.Verify(Authorization)
		if verifyErr != nil {
			return c.JSON(401, "Error on Sign")
		}
		if !verify {
			return c.JSON(401, "Error on Token")
		}

		userId, subErr := m.extractor.ExtractSubject(Authorization)
		if subErr != nil || userId == "" {
			return c.JSON(401, "Error on Token")
		}

		permission, aclErr := m.acl.CheckPermission(userId, c.Path(), c.Request().Method)
		if aclErr != nil {
			jsonErr := c.JSON(401, "no permission")
			return errors.Join(jsonErr, aclErr)
		}

		if !permission {
			return c.JSON(401, "no permission")
		}

		return next(c)
	}
}
