package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectPostgresGORM_Success(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Nil(t, err)

	db, err := ConnectPostgresGORM()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
