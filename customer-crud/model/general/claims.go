package general

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type Claim struct {
	ID uuid.UUID
	jwt.StandardClaims
}
