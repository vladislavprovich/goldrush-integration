package storage

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRedis(t *testing.T) (*Redis, *miniredis.Miniredis, func()) {
	// Start a mini Redis server for testing
	mr, err := miniredis.Run()
	require.NoError(t, err)

	// Create a Redis client connected to the mini Redis server
	cfg := &Config{
		Addr:     mr.Addr(),
		Password: "",
		DB:       0,
	}

	redis := New(cfg)

	// Return cleanup function
	cleanup := func() {
		mr.Close()
	}

	return redis, mr, cleanup
}

func TestRedis_TokenOperations(t *testing.T) {
	ctx := context.Background()
	redis, _, cleanup := setupTestRedis(t)
	defer cleanup()

	// Test data
	email := "test@example.com"
	token := "test-token-12345"
	expiration := time.Hour

	// Test SaveToken
	err := redis.SaveToken(ctx, email, token, expiration)
	assert.NoError(t, err)

	// Test HasToken
	exists, err := redis.HasToken(ctx, email)
	assert.NoError(t, err)
	assert.True(t, exists)

	// Test GetToken
	retrievedToken, err := redis.GetToken(ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, token, retrievedToken)

	// Test DeleteToken
	err = redis.DeleteToken(ctx, email)
	assert.NoError(t, err)

	// Verify token is deleted
	exists, err = redis.HasToken(ctx, email)
	assert.NoError(t, err)
	assert.False(t, exists)

	// Verify GetToken returns error for non-existent token
	_, err = redis.GetToken(ctx, email)
	assert.Error(t, err)
}

func TestRedis_GetToken_NonExistent(t *testing.T) {
	ctx := context.Background()
	redis, _, cleanup := setupTestRedis(t)
	defer cleanup()

	// Test getting a non-existent token
	_, err := redis.GetToken(ctx, "nonexistent@example.com")
	assert.Error(t, err)
}

func TestRedis_TokenExpiration(t *testing.T) {
	ctx := context.Background()
	redis, mr, cleanup := setupTestRedis(t)
	defer cleanup()

	// Test data
	email := "expiration@example.com"
	token := "expiring-token"
	expiration := time.Second

	// Save token with short expiration
	err := redis.SaveToken(ctx, email, token, expiration)
	assert.NoError(t, err)

	// Verify token exists
	exists, err := redis.HasToken(ctx, email)
	assert.NoError(t, err)
	assert.True(t, exists)

	// Fast-forward time in miniredis
	mr.FastForward(2 * time.Second)

	// Verify token no longer exists after expiration
	exists, err = redis.HasToken(ctx, email)
	assert.NoError(t, err)
	assert.False(t, exists)
}
