package authservice

import (
	"strings"
	"taskema/pkg/richerror"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	ContextKey            string        `koanf:"context_key"`
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Service struct {
	cfg Config
}

func New(cfg Config) Service {
	return Service{cfg: cfg}
}

func (s Service) GenerateAccessToken(userID uint, roleID uint) (string, error) {
	return s.generateToken(userID, roleID, s.cfg.AccessSubject, s.cfg.AccessExpirationTime)
}

func (s Service) GenerateRefreshToken(userID uint, roleID uint) (string, error) {
	return s.generateToken(userID, roleID, s.cfg.RefreshSubject, s.cfg.RefreshExpirationTime)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	op := "authservice.ParseToken"

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.SignKey), nil
	})
	if err != nil {
		return nil,
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil,
			richerror.New(op).
				WithMessage(richerror.MsgErrorCastingClaims).
				WithCode(richerror.CodeUnexpected)
	}
}

func (s Service) generateToken(userID uint, roleID uint, subject string, expireDuration time.Duration) (string, error) {
	op := "authservice.generateToken"

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
		RoleID: roleID,
	}

	// TODO - add sign method to config
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.cfg.SignKey))
	if err != nil {
		return "",
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	return tokenString, nil
}
