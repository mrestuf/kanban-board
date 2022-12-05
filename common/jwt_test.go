package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mrestuf/kanban-board/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateToken_Success(t *testing.T) {
	config.GetJwtSignature = func() []byte {
		return []byte("123test")
	}

	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiZXhwIjoxOTY4NTc5NDk0LCJpc3MiOiIxMzEyIiwic3ViIjoibmF0aGFuMTIzQGdtYWlsLmNvbSJ9.5oGlxgEIYdWkCbM8uI5Sv0UqiD1eEhEXtiKarUAfTNY"

	claims, err := ValidateToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, 3, claims.Id)
	assert.Equal(t, "nathan123@gmail.com", claims.Subject)
}

func TestValidateToken_InvalidSignature(t *testing.T) {
	config.GetJwtSignature = func() []byte {
		return []byte("123test")
	}

	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiZXhwIjoxOTY4NTc5NDk0LCJpc3MiOiIxMzEyIiwic3ViIjoibmF0aGFuMTIzQGdtYWlsLmNvbSJ9.-Cx70LKZgPI67IzPKHGPgv_F3w0IV8klkpyAr35Gh-c"

	claims, err := ValidateToken(tokenString)
	assert.Nil(t, claims)
	assert.NotNil(t, err)
	assert.Equal(t, jwt.ErrSignatureInvalid, err.(*jwt.ValidationError).Inner)
}
