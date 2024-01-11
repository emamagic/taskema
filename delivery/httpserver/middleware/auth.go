package middleware

import (
	"taskema/service/authservice"

	mw "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mw.WithConfig(mw.Config{
		ContextKey: config.ContextKey,
		SigningKey: []byte(config.SignKey),
		// TODO - as sign method string to config...
		SigningMethod: "HS256",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}

			return claims, nil
		},
	})
}
