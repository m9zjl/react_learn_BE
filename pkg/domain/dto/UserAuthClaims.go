package dto

import (
	"github.com/golang-jwt/jwt"
	"server/pkg/domain/entity"
)

type UserAuthClaim struct {
	jwt.StandardClaims
	entity.User
}
